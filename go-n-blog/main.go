package main

import (
	"bufio"
	"context"
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
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Printf("Server online. Listening on %s\n", server.Addr)
	go server.ListenAndServe()
	for {
		s := bufio.NewScanner(os.Stdin)
		fmt.Print(`Welcome
	Press "q" to kill the server
	Press "t" to force templates reloading
>: `)

		s.Scan()
		switch s.Text() {
		case "q":
			fmt.Printf("Closing server...\n")
			err := server.Shutdown(context.TODO())
			fmt.Printf("Server closed. %s", err)
			return
		case "t":
			fmt.Println("Realoading templates...")
			tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
			fmt.Println("Templates reloaded.")
		default:
			fmt.Println("Input not valid.")
		}

	}
}

func home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in index template! %s\n", err)
	}
}
