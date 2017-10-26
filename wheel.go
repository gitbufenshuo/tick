package tick

type Wheel struct {
	nowTick   uint8
	slotNum   uint8
	evListArr []*EventLinkList
	God       *TickWheel
}

func NewWheel(slotNum uint8, god *TickWheel) *Wheel {
	wheel := Wheel{
		slotNum: slotNum,
		God:     god,
	}
	wheel.evListArr = make([]*EventLinkList, slotNum, slotNum)
	for idx := range wheel.evListArr {
		wheel.evListArr[idx] = NewEventLinkList()
	}
	return &wheel
}
func (wheel *Wheel) AddEventAt(slot uint8, ev *Event) {
	wheel.evListArr[slot].ADD(ev)
}

// 溢出了么
func (wheel *Wheel) Tock() bool {
	var flow bool
	wheel.nowTick += 1
	if wheel.nowTick == wheel.slotNum {
		flow = true
		wheel.nowTick = 0
	}
	// 遍历自己的相应 slot 处理所有事件
	evList := wheel.evListArr[wheel.nowTick]
	var dealwith = func(ev *Event) {
		wheel.God.Do(ev, wheel.God)
	}
	evList.Scan(dealwith)
	return flow
}
