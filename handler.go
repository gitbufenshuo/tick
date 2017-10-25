package tick

type Handlers struct {
	handlerMap map[uint32]func(*Event)
}

func NewHandlers() *Handlers {
	var handlers Handlers
	handlers.handlerMap = make(map[uint32]func(*Event))
	return &handlers
}
func (handlers *Handlers) AddHandler(eid uint32, handler func(*Event)) {
	m := handlers.handlerMap
	if _, ok := m[eid]; ok {
		panic("duplicate event id")
	}
	m[eid] = handler
}

func AddEvent(ev *Event) {

}
