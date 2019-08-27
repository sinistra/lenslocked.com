package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var port = ":3001"

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Frequently asked questions.....")
}

func fourofour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you "+
		"were looking for :(</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	var h http.Handler = http.HandlerFunc(fourofour)
	r.NotFoundHandler = h
	http.ListenAndServe(port, r)
}
