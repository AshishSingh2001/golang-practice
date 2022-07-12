package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("mid gs", runtime.NumGoroutine())
	fmt.Println("CPU :", runtime.NumCPU())
	fmt.Println("arch :", runtime.GOARCH)
	fmt.Println("OS :", runtime.GOOS)
}
