package main

// import (
// 	"fmt"
// )

// type person struct {
// 	first string
// 	last  string
// }

// func (p *person) speak() {
// 	fmt.Println(*p)
// }

// type human interface {
// 	speak()
// }

// func saySomething(h human) {
// 	h.speak()
// }

// func main() {

// 	p := person{
// 		"Hello", "paaji",
// 	}
// 	// This gives an error as
// 	// Receivers       Values
// 	// -----------------------------------------------
// 	// (t T)           T and *T
// 	// (t *T)          *T
// 	// saySomething(p)
	
// 	saySomething(&p)
// }
