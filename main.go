package main

import (
	"fmt"
	"net/http"
	//"github.com/julienschmidt/httprouter"
)

// function call answer request user

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//fmt.Fprintf(w, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to my awesome site. My  friend!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep  being sent to an invalid page.</p>)		")
	}

	// fmt.Fprintf(w, "<h1>Welcome to my awesome site. My  friend!</h1>")
	// fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func main() {
	//mux := &http.ServeMux{}
	//mux.HandleFunc("/", handlerFunc)
	//r := mux.NewRouter()
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
	//router := httprouter.New()
	//router.GET("/hello/:name", Hello)

}
