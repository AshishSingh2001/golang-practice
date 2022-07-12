package main

import "fmt"

type hotdog int

var a hotdog
var y int

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	a = 30
	fmt.Println(a)
	y = int(a)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
