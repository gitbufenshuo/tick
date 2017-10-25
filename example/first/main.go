package main

import "github.com/gitbufenshuo/tick"
import "fmt"

func main() {
	evlist := tick.NewEventLinkList()
	ev1 := tick.NewEvent(1, "helloworld")
	ev2 := tick.NewEvent(2, "aaaaaaaaaa")
	ev3 := tick.NewEvent(3, "bbbbbbbbbb")
	ev4 := tick.NewEvent(4, "cccccccccc")
	n1 := evlist.ADD(ev1)
	evlist.Scan()
	n2 := evlist.ADD(ev2)
	evlist.Scan()
	evlist.ADD(ev3)
	evlist.Scan()
	n4 := evlist.ADD(ev4)
	evlist.Scan()

	//
	fmt.Println("===")
	evlist.Remove(n1)
	evlist.Scan()
	evlist.Remove(n2)
	evlist.Scan()
	evlist.Remove(n4)
	evlist.Scan()
}
