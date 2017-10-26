package tick

import (
	"fmt"
)

type TickWheel struct {
	nowTick    uint64
	levelWheel []*Wheel
	*Handlers
}

func NewTickWheel() *TickWheel {
	var tickWheel TickWheel
	levelWheel := make([]*Wheel, 6, 6)
	levelWheel[0] = NewWheel(200, &tickWheel) // [0   -> 199]                     199
	levelWheel[1] = NewWheel(100, &tickWheel) // [200 -> 19999]                   99*200 + 199 = 19999
	levelWheel[2] = NewWheel(100, &tickWheel) // [20000 -> 1999999]               99*20000 + 19999 = 1999999
	levelWheel[3] = NewWheel(100, &tickWheel) // [2000000 -> 199999999]           99*2000000 + 1999999 = 199999999
	levelWheel[4] = NewWheel(50, &tickWheel)  // [200000000 -> 9999999999]        49*200000000 + 199999999 = 9999999999
	levelWheel[5] = NewWheel(50, &tickWheel)  // [10000000000 -> 499999999999]    49*10000000000 +  9999999999 = 499999999999
	tickWheel.levelWheel = levelWheel
	tickWheel.Handlers = NewHandlers()
	return &tickWheel
}
func (tickWheel *TickWheel) SetTick(tick uint64) {
	if tick > 499999999999 {
		panic("oh no 499999999999")
	}
	tickWheel.nowTick = tick
	for idx, slot := range tickWheel.SlotsAt(tick) {
		tickWheel.levelWheel[idx].nowTick = slot
	}
}
func (tickWheel *TickWheel) SlotsAt(tick uint64) []uint8 {
	if tick > 499999999999 {
		panic("oh no 499999999999")
	}
	res := make([]uint8, 6, 6)

	res[0] = uint8(tick % 200)
	tick = tick / 200

	res[1] = uint8(tick % 100)
	tick = tick / 100

	res[2] = uint8(tick % 100)
	tick = tick / 100

	res[3] = uint8(tick % 100)
	tick = tick / 100

	res[4] = uint8(tick % 50)
	tick = tick / 50

	res[5] = uint8(tick % 50)
	tick = tick / 50
	return res
}
func (tickWheel *TickWheel) ShowTick() {
	fmt.Printf("%v ->", tickWheel.nowTick)
	for idx := range tickWheel.levelWheel {
		fmt.Printf(" %v", tickWheel.levelWheel[idx].nowTick)
	}
	fmt.Println()
}
func (tickWheel *TickWheel) AddEventAt(absSlot *AbsSlot, ev *Event) (*EventLinkList, *EventNode) {
	enode := tickWheel.levelWheel[absSlot.Wheel].evListArr[absSlot.Slot].ADD(ev)
	return tickWheel.levelWheel[absSlot.Wheel].evListArr[absSlot.Slot], enode
}

func (tickWheel *TickWheel) RegisterEventAfter(tick uint64, ev *Event) {
	if ev.Eid == 0 {
		panic(" 注册自定义事件时，事件 id 不能为 0 ")
	}
	if tick == 0 {
		return
	}
	fmt.Printf("REGISTER now:[%v] after:[%v]\n", tickWheel.nowTick, tick)
	lastTick := tickWheel.nowTick + tick
	lastTickSlots := tickWheel.SlotsAt(lastTick)
	var absSlots []*AbsSlot
	var highLevel int
	for idx := 5; idx != -1; idx-- {
		if tickWheel.levelWheel[idx].nowTick != lastTickSlots[idx] {
			highLevel = idx
			absSlots = make([]*AbsSlot, 0, idx+1)
			newAbsSlot := new(AbsSlot)
			newAbsSlot.Wheel = highLevel
			newAbsSlot.Slot = lastTickSlots[idx]
			absSlots = append(absSlots, newAbsSlot)
			break
		}
	}
	for idx := highLevel - 1; idx != -1; idx-- {
		if lastTickSlots[idx] != 0 {
			newAbsSlot := new(AbsSlot)
			newAbsSlot.Wheel = idx
			newAbsSlot.Slot = lastTickSlots[idx]
			absSlots = append(absSlots, newAbsSlot)
		}
	}
	for _, absslot := range absSlots {
		fmt.Println("{}{}{}|||| ", absslot.Wheel, absslot.Slot)
	}
	newResetData := new(ResetData)
	newResetData.NowIdx = 0
	newResetData.RealEvent = ev
	newResetData.Slots = absSlots
	newResetEvent := NewEvent(0, newResetData)
	tickWheel.AddEventAt(absSlots[0], newResetEvent)
}
func (tickWheel *TickWheel) Tock() {
	// 从低层到高层，遍历 tick wheel，检查所有事件，并处理
	tickWheel.nowTick += 1
	var flow bool
	for _, wheel := range tickWheel.levelWheel {
		if flow = wheel.Tock(); !flow {
			break
		}
	}
}

type AbsSlot struct {
	Wheel int
	Slot  uint8
}
