package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("dog.gohtml"))
}

func main() {

	s, _ := os.Getwd()
	fmt.Println(s)

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/duck.jpeg", img)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Foo Ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", "Hello from dog")
}

func img(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "duck.jpeg")
}
