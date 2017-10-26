package main

import "github.com/gitbufenshuo/tick"

func testTickWheel() {
	tickWheel := tick.NewTickWheel()
	tickWheel.AddHandler(1, new(SelfPrintEventHandler))

	helloWorldData := new(SelfPrintEvent)
	helloWorldData.SelfInfo = "hello world"
	helloworldevent := tick.NewEvent(1, helloWorldData)
	tickWheel.RegisterEventAfter(33333, helloworldevent)

	for idx := 0; idx != 33339; idx++ {
		tickWheel.Tock()
		if idx == 2222 {
			hahaData := new(SelfPrintEvent)
			hahaData.SelfInfo = "haha"
			hahaEvent := tick.NewEvent(1, hahaData)
			tickWheel.RegisterEventAfter(31111, hahaEvent)
		}
	}
}
func main() {
	testTickWheel()
}
