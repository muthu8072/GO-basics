package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type muth struct {
	Title   string
	Heading string
	Value   string
}

func main() {
	myHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("template.html")
		//data := "hello wel"
		p := muth{Title: "great hari", Heading: "hii boys", Value: "worth"}
		t.Execute(w, p)

	}
	http.HandleFunc("/", myHandlerFunc)

	//	http.HandleFunc("/cat",mycatHadlerfunc)

	http.ListenAndServe(":9000", nil)
}
func mycatHandlerfunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "cat.html")
	} else {
		user := r.FormValue("username")
		pass := r.FormValue("password")
		fmt.Fprintf(w, "username is %s and password is %s", user, pass)
	}
}
