package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", pageHandlerFactory(homeView))
	r.HandleFunc("/contact", pageHandlerFactory(contactView))
	r.HandleFunc("/faq", pageHandlerFactory(faqView))

	http.ListenAndServe(":3000", r)
}

func pageHandlerFactory(v *views.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := v.Render(w, nil)
		if err != nil {
			panic(err)
		}
	}
}
