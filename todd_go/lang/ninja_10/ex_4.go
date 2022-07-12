package main

// import (
// 	"fmt"
// )

// func main() {
// 	q := make(chan int)
// 	c := gen(q)

// 	receive(c, q)

// 	fmt.Println("about to exit")
// }

// func gen(q chan<- int) <-chan int {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 		q <- 999
// 		close(c)
// 	}()
// 	return c
// }

// func receive(c, q <-chan int) {
// 	for {
// 		select {
// 		case ci := <-c:
// 			fmt.Println("from c :", ci)
// 		case i := <-q:
// 			fmt.Println("from q :", i)
// 			if i != 0 {
// 				return
// 			}
// 		}
// 	}
// }
