package main

import (
	"html/template"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w,"index.gohtml","Hemlo from the root")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w,"index.gohtml","Good Morning")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w,"index.gohtml","woof woof")
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("/media/lancelot/hard_drive/Coding/golang/todd_go/web/web_3/template/*"))
}

func main() {
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)

}
