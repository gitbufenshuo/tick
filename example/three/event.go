package main

import (
	"fmt"

	"github.com/gitbufenshuo/tick"
	"github.com/gitbufenshuo/tick/example/three/entity"
)

type SelfPrintEventHandler struct {
}
type SelfPrintEvent struct {
	SelfInfo string
}

func (selfPrintEventHandler *SelfPrintEventHandler) Do(ev *tick.Event, god *tick.TickWheel) {
	fmt.Println("SELF EVENT SUCCESS " + ev.Data.(*SelfPrintEvent).SelfInfo)
}

type PeopleMoveEventHandler struct {
}
type PeopleMoveData struct {
	*entity.People
	TargetPos int
}

func (peopleMoveEventHandler *PeopleMoveEventHandler) Do(ev *tick.Event, god *tick.TickWheel) {
	data := ev.Data.(*PeopleMoveData)
	if data.Pos == data.TargetPos {
		return
	}
	displacement := data.TargetPos - data.Pos%2
	data.Move(displacement)

	nextMove := tick.NewEvent(2, data)

	god.RegisterEventAfter(1, nextMove)
}
