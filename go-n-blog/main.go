package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

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
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in index template! %s", err)
	}

}
