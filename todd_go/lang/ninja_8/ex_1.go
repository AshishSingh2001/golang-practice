package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type User struct {
// 	First string
// 	Age   int
// }

// func main() {
// 	u1 := User{
// 		First: "James",
// 		Age:   32,
// 	}

// 	u2 := User{
// 		First: "Moneypenny",
// 		Age:   27,
// 	}

// 	u3 := User{
// 		First: "M",
// 		Age:   54,
// 	}

// 	users := []User{u1, u2, u3}

// 	fmt.Println(users)

// 	bs, err := json.Marshal([]User{u1, u2, u3})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	fmt.Println(string(bs))
// 	// your code goes here
// }
