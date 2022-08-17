package main

import (
	"fmt"
	"net/http"
)

// function call answer request user

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//fmt.Fprintf(w, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to my awesome site. My  friend!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	}

	// fmt.Fprintf(w, "<h1>Welcome to my awesome site. My  friend!</h1>")
	// fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
