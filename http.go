package main

import (
	"fmt"
	"net/http"
)

func main() {
	myHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		fmt.Fprintf(w, "<h1>hello muthuraja</h1>")
	}
	http.HandleFunc("/", myHandlerFunc)
	http.HandleFunc("/ca", myHandlerfunc)
	http.ListenAndServe(":9000", nil)
}
func myHandlerfunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello cat")
	http.ServeFile(w, r, "cat.html") //html file name
}
