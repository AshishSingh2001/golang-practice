package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	count := 3
// 	wg.Add(count)

// 	var ctr int = 0
// 	for i := 0; i < count; i++ {
// 		go func() {
// 			c := ctr
// 			time.Sleep(1 * time.Second)
// 			// runtime.Gosched()
// 			c++
// 			ctr = c
// 			fmt.Println(ctr)
// 			wg.Done()
// 		}()
// 	}

// 	fmt.Println("mid gs", runtime.NumGoroutine())
// 	wg.Wait()
// 	fmt.Println(ctr)
// 	fmt.Println("about to exit")
// 	fmt.Println("end gs", runtime.NumGoroutine())
// }
