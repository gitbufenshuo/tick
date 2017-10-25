package main

import "github.com/gitbufenshuo/tick"

func main() {
	tickWheel := tick.NewTickWheel()
	tickWheel.ShowTick()
	tickWheel.SetTick(201)
	tickWheel.ShowTick()

}
