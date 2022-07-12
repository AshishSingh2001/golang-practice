package main

import "fmt"

func foo(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	x := foo(3)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	y := foo(-3)
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}
