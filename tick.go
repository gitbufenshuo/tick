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

	levelWheel := make([]*Wheel, 6, 6)
	levelWheel[1] = NewWheel(200) // [0   -> 199]                     199
	levelWheel[0] = NewWheel(100) // [200 -> 19999]                   99*200 + 199 = 19999
	levelWheel[2] = NewWheel(100) // [20000 -> 1999999]               99*20000 + 19999 = 1999999
	levelWheel[3] = NewWheel(100) // [2000000 -> 199999999]           99*2000000 + 1999999 = 199999999
	levelWheel[4] = NewWheel(50)  // [200000000 -> 9999999999]        49*200000000 + 199999999 = 9999999999
	levelWheel[5] = NewWheel(50)  // [10000000000 -> 499999999999]    49*10000000000 +  9999999999 = 499999999999
	return &TickWheel{
		levelWheel: levelWheel,
	}
}

func (tickWheel *TickWheel) SetTick(tick uint64) {
	if tick > 499999999999 {
		panic("oh no 499999999999")
	}
	tickWheel.nowTick = tick
	tickWheel.levelWheel[0].nowTick = uint8(tick % 200)
	tick = tick / 200

	tickWheel.levelWheel[1].nowTick = uint8(tick % 100)
	tick = tick / 100

	tickWheel.levelWheel[2].nowTick = uint8(tick % 100)
	tick = tick / 100

	tickWheel.levelWheel[3].nowTick = uint8(tick % 100)
	tick = tick / 100

	tickWheel.levelWheel[4].nowTick = uint8(tick % 50)
	tick = tick / 50

	tickWheel.levelWheel[5].nowTick = uint8(tick % 50)
	tick = tick / 50
}
func (tickWheel *TickWheel) ShowTick() {
	fmt.Printf("%v ->", tickWheel.nowTick)
	for idx := range tickWheel.levelWheel {
		fmt.Printf(" %v", tickWheel.levelWheel[idx].nowTick)
	}
	fmt.Println()
}
func (tickWheel *TickWheel) AddEventAt(wheel int, slot uint8, ev *Event) (*EventLinkList, *EventNode) {
	enode := tickWheel.levelWheel[wheel].evListArr[slot].ADD(ev)
	return tickWheel.levelWheel[wheel].evListArr[slot], enode
}
