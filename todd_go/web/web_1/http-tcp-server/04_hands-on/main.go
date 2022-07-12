package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	// read req
	url := request(conn)

	if strings.HasSuffix(url, "apply") {
		apply(conn)
	} else {
		// write res
		respond(conn)

	}

}

func request(conn net.Conn) string {
	n := 0
	scanner := bufio.NewScanner(conn)
	var url string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if n == 0 {
			m := strings.Fields(ln)
			url = m[1]
			fmt.Println("***Method", m[0], "***Url", m[1], "Http version", m[2])
		}
		if ln == "" {
			// headers are done
			break
		}
		n++
	}
	return url
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Welcome To Apply</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
