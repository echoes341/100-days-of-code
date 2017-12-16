package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	mux = httprouter.New()
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
	initRoutes(mux)
}

func main() {
	http.ListenAndServe(":8080", mux)
}

func home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
