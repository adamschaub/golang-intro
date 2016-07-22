package main

import "fmt"

type Me struct {
	Adam   bool
	Schaub bool
}

func (me Me) String() string {
	if me.Adam == true && me.Schaub == true {
		return "Adam Schaub"
	}
	return "Some poser"
}

func main() {
	m := Me{Adam: true, Schaub: true}
	fmt.Println(m)
}
