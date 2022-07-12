package main

// import (
// 	"fmt"
// )

// func main() {
// 	c := gen()
// 	receive(c)
// 	fmt.Println("about to exit")
// }

// func gen() <-chan int {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 	}()

// 	return c
// }

// func receive(c <-chan int) {
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }
