package main

import "fmt"

type test interface {
	changeMe(string)
}

type person struct {
	name string
}

func (p *person) changeMe(newName string) {
	fmt.Printf("%T\n", p.name)
	p.name = newName
	(*p).name = "this shouldnt change"
}

func info(t test) {
	fmt.Printf("%T\n", t)
	fmt.Println("before ", t)
	t.changeMe("NewName")
	fmt.Println("after", t)
}

func main() {
	p := &person{
		name: "Jason",
	}
	info(p)

}
