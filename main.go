package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sinistra/lenslocked.com/controllers"
)

var port = ":3001"

//A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func fourofour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you "+
		"were looking for :(</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	var h http.Handler = http.HandlerFunc(fourofour)
	r.NotFoundHandler = h

	fmt.Println("Server running on port " + port)
	http.ListenAndServe(port, r)
}
