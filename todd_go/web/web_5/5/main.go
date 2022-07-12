package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public/pic"))
	http.HandleFunc("/", foo)
	http.Handle("/public/pic/", http.StripPrefix("/public/pic", fs))
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	if err := tpl.Execute(w, nil); err != nil {
		log.Fatalln("Template didnt execute", err)
	}
}
