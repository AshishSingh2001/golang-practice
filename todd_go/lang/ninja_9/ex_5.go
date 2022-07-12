package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"sync/atomic"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	count := 100
// 	wg.Add(count)
// 	var ctr int64
// 	for i := 0; i < count; i++ {
// 		go func() {
// 			fmt.Println("ctr :", atomic.AddInt64(&ctr, 1))
// 			wg.Done()
// 		}()
// 	}

// 	// fmt.Println("mid gs", runtime.NumGoroutine())
// 	wg.Wait()
// 	fmt.Println(ctr)
// 	fmt.Println("about to exit")
// 	fmt.Println("end gs", runtime.NumGoroutine())
// }
