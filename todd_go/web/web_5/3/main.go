package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/",foo)
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w)
	if err :=  tpl.Execute(w,nil); err != nil {
		log.Fatalln("Template didnt execute",err)
	}
}
