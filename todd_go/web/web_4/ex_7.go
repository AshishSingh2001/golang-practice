package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"strings"
// )

// func main() {
// 	lis, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer lis.Close()

// 	for {
// 		conn, err := lis.Accept()
// 		if err != nil {
// 			fmt.Println("concern")
// 			log.Fatalln(err)
// 		}
// 		go serve(conn)
// 	}
// }

// func serve(conn net.Conn) {
// 	defer conn.Close()
// 	sc := bufio.NewScanner(conn)
// 	x := 0
// 	for sc.Scan() {
// 		ln := sc.Text()
// 		if ln == "" {
// 			break
// 		}
// 		if x == 0 {
// 			x++
// 			s := strings.Split(ln, " ")
// 			fmt.Println("Request Method : ", s[0])
// 			fmt.Println("Request URL : ", s[1])
// 			fmt.Println("HTTP Version : ", s[2])
// 		} else {
// 			fmt.Println(ln)
// 		}
// 	}
// 	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

// 	body := "THis is super cool"

// 	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

// 	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
// 	io.WriteString(conn, "\r\n")
// 	io.WriteString(conn, body)
// }
