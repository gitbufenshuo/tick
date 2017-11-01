package tick

type Handler interface {
	Do(*Event, *TickWheel)
}
type Handlers struct {
	handlerMap map[uint32]Handler
	god        *TickWheel
}

func NewHandlers(god *TickWheel) *Handlers {
	var handlers Handlers
	handlers.handlerMap = make(map[uint32]Handler)
	handlers.AddHandler(0, new(ResetEventHandler)) // 内置 0 号 ， 重设事件 handler
	handlers.god = god
	return &handlers
}
func (handlers *Handlers) AddHandler(eid uint32, handler Handler) {
	m := handlers.handlerMap
	if _, ok := m[eid]; ok {
		panic("duplicate event id")
	}
	m[eid] = handler
}
func (handlers *Handlers) Do(ev *Event) {
	m := handlers.handlerMap
	if handler, ok := m[ev.Eid]; !ok {
		panic("event 包含了不识别的 id")
	} else {
		handler.Do(ev, handlers.god)
	}
}

type ResetEventHandler struct {
}
type ResetData struct {
	Slots     []*AbsSlot // 高层 wheel 排在前面
	NowIdx    int
	RealEvent *Event
}

func (ResetEventHandler *ResetEventHandler) Do(ev *Event, god *TickWheel) {
	data := ev.Data.(*ResetData)
	if data.NowIdx == len(data.Slots)-1 {
		god.Do(data.RealEvent)
		return
	}
	data.NowIdx += 1
	newResetEvent := NewEvent(0, data)
	god.AddEventAt(data.Slots[data.NowIdx], newResetEvent)
}
