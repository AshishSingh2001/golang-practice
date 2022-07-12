package main

// import (
// 	"fmt"
// )

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 4
// 		c <- 5
// 		close(c)
// 	}()

// 	for i := 0; i < 4; i++ {
// 		v, ok := <-c
// 		fmt.Println(v, ok)
// 	}
// }
