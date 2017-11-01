package entity

import (
	"fmt"
)

type People struct {
	Name string
	Pos  int
}

func (people *People) Move(displacement int) {
	people.Pos += displacement
	fmt.Printf("[%v] moves [%v]: now at [%v]\n", people.Name, displacement, people.Pos)
}
