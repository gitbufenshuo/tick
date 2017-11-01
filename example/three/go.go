package main

import "github.com/gitbufenshuo/tick"
import "github.com/gitbufenshuo/tick/example/three/entity"

var littleMing entity.People

func playerinit() {
	littleMing.Name = "小明"
	littleMing.Pos = 0
}
func testTickWheel() {
	playerinit()
	tickWheel := tick.NewTickWheel()
	tickWheel.AddHandler(1, new(SelfPrintEventHandler))
	tickWheel.AddHandler(2, new(PeopleMoveEventHandler))

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
		if idx == 33222 {
			moveData := new(PeopleMoveData)
			moveData.People = &littleMing
			moveData.TargetPos = 100
			moveEvent := tick.NewEvent(2, moveData)
			tickWheel.RegisterEventAfter(1, moveEvent)
		}
	}
}
func main() {
	testTickWheel()
}
