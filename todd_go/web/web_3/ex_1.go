package main

// import (
// 	"io"
// 	"net/http"
// )

// func root(w http.ResponseWriter, req *http.Request)  {
// 	io.WriteString(w,"Hemlo from the root")
// }

// func me(w http.ResponseWriter, req *http.Request)  {
// 	io.WriteString(w,"Good morning")
// }

// func dog(w http.ResponseWriter, req *http.Request)  {
// 	io.WriteString(w,"woof woof")
// }

// func main() {
// 	http.HandleFunc("/",root);
// 	http.HandleFunc("/dog/",dog);
// 	http.HandleFunc("/me/",me);

// 	http.ListenAndServe(":8080",nil);

// }