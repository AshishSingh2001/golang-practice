package main

// import (
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
// 		print(conn.LocalAddr().String())
// 		io.WriteString(conn,"I see you connected\n")
// 		conn.Close()
// 		if err != nil {
// 			log.Fatalln(err)
// 		} else {
// 			break;
// 		}
// 	}

// }