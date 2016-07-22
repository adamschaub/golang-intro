package main

import "fmt"

type Person struct{ Name string }
type Employee struct {
	Person
	Company string
}

func (p Person) SayHi() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func main() {
	emp := Employee{Person{"Matt Coughfield"}, "Cisgoat Systems"}
	emp.SayHi()
}
