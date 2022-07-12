package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	count := 3
// 	wg.Add(count)
// 	mu := sync.Mutex{}
// 	var ctr int = 0
// 	for i := 0; i < count; i++ {
// 		go func() {
// 			mu.Lock()
// 			c := ctr
// 			c++
// 			ctr = c
// 			fmt.Println("ctr :", ctr)
// 			mu.Unlock()
// 			wg.Done()
// 		}()
// 	}

// 	// fmt.Println("mid gs", runtime.NumGoroutine())
// 	wg.Wait()
// 	fmt.Println(ctr)
// 	fmt.Println("about to exit")
// 	fmt.Println("end gs", runtime.NumGoroutine())
// }
