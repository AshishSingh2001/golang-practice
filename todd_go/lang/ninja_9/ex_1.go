package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {

// 	fmt.Println("begin cpu", runtime.NumCPU())
// 	fmt.Println("begin gs", runtime.NumGoroutine())

// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		fmt.Println("Goroutine 2")
// 		wg.Done()
// 	}()

// 	go func() {
// 		fmt.Println("Gorutine 3")
// 		wg.Done()
// 	}()
// 	fmt.Println("mid gs", runtime.NumGoroutine())

// 	wg.Wait()

// 	fmt.Println("about to exit")
// 	fmt.Println("end gs", runtime.NumGoroutine())
// }
