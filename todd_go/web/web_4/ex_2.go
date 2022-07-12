package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// )

// func main() {
// 	lis,err := net.Listen("tcp", ":8080");
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer lis.Close()
	
// 	for {
// 		conn,err := lis.Accept()
// 		if err != nil {
// 			fmt.Println("concern")
// 			log.Fatalln(err)
// 		}
// 		sc := bufio.NewScanner(conn)
// 		for sc.Scan() {
// 			ln := sc.Text()
// 			fmt.Println(ln)
// 		}
// 		fmt.Println("Code got here.")
// 		_,err = io.WriteString(conn, "I see you connected.")
// 		if err != nil {
// 			log.Fatalln(err)
// 			} else {
// 				break;
// 			}
// 		conn.Close()
// 	}

// }