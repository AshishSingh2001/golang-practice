package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	// q := make(chan bool)
	wg := sync.WaitGroup{}

	// go func() {
	// 	for j := 0; j < 10; j++ {
	// 		wg.Add(j)
	// 		go func() {
	// 			for i := 0; i < 10; i++ {
	// 				c <- i
	// 			}
	// 			wg.Done()
	// 		}()
	// 	}
	// 	wg.Wait()
	// 	close(c)
	// }()

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for i := 0; i < 10; i++ {
					c <- i*10 + m
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
