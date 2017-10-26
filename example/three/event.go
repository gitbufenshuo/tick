package main

import (
	"fmt"

	"github.com/gitbufenshuo/tick"
)

type SelfPrintEventHandler struct {
}

type SelfPrintEvent struct {
	SelfInfo string
}

func (selfPrintEventHandler *SelfPrintEventHandler) Do(ev *tick.Event, god *tick.TickWheel) {
	fmt.Println("SELF EVENT SUCCESS " + ev.Data.(*SelfPrintEvent).SelfInfo)
}
