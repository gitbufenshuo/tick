package tick

type Event struct {
	Eid  uint32
	Data interface{}
	Res  int
	Tick uint64
}

func NewEvent(eid uint32, data interface{}) *Event {
	return &Event{
		eid,
		data,
		0,
		0,
	}
}

type EventNode struct {
	prev *EventNode
	next *EventNode
	curr *Event
}

type EventLinkList struct {
	header *EventNode
	tail   *EventNode
	elen   uint32
}

func NewEventLinkList() *EventLinkList {
	return &EventLinkList{}
}
func (ell *EventLinkList) LEN() uint32 {
	return ell.elen
}
func (ell *EventLinkList) ADD(enode *Event) *EventNode {
	newnode := new(EventNode)
	newnode.curr = enode
	if ell.elen == 0 {
		ell.header = newnode
	} else {
		newnode.prev = ell.tail
		ell.tail.next = newnode
	}
	ell.tail = newnode
	ell.elen += 1
	return newnode
}
func (ell *EventLinkList) Scan(dealwith func(ev *Event)) {
	if ell.elen == 0 {
		return
	}
	var tenode = ell.header

	for tenode != nil {
		dealwith(tenode.curr)
		tenode = tenode.next
	}
}
func (ell *EventLinkList) Remove(enode *EventNode) {
	if ell.header == enode {
		if ell.tail == enode {
			ell.header = nil
			ell.tail = nil
			ell.elen = 0
			return
		} else {
			ell.header = enode.next
			ell.header.prev = nil
			ell.elen -= 1
			return
		}
	}
	if ell.tail == enode {
		ell.elen -= 1
		ell.tail = enode.prev
		ell.tail.next = nil
	} else {
		ell.elen -= 1
		enode.prev.next = enode.next
		enode.next.prev = enode.prev
	}
}
