package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() string {
	return fmt.Sprintln("uptown, func you up, uptown func you up")
}

func (sa secretAgent) speak() string {
	return fmt.Sprintln("shaken, not stirred")
}

type humanoid interface {
	speak() string
}

func vomit(h humanoid) {
	fmt.Printf("%T\n", h)
	fmt.Println(h)
	switch v := h.(type) {
	case person:
		fmt.Println(v.fName)
	case secretAgent:
		fmt.Println(v.person)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	p1 := person{"Nina", "Simone", 25}
	sa1 := secretAgent{person{"Ian", "Fleming", 42}, false}
	vomit(p1)
	vomit(sa1)
}
