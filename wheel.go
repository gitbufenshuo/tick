package tick

type Wheel struct {
	nowTick   uint8
	slotNum   uint8
	evListArr []*EventLinkList
}

func NewWheel(slotNum uint8) *Wheel {
	wheel := Wheel{
		slotNum: slotNum,
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
