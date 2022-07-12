package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("counter")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.SetCookie(w, &http.Cookie{
				Name:  "counter",
				Value: "1",
			})
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		return
	}
	ctr, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "counter",
		Value: fmt.Sprint(ctr + 1),
	})
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
