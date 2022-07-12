package main

import "fmt"

func main() {
	x := struct{
		isAnon bool
		name string
	} {
		isAnon: true,
		name: "Anonymous Struct",
	}

	fmt.Println(x)
}
