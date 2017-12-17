package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" //for setup
	"github.com/julienschmidt/httprouter"
)

const (
	menuStr = `Welcome
	Press "q" to kill the server
	Press "t" to force templates reloading
>: `
)

var (
	mux = httprouter.New()
	tpl *template.Template
	db  *sql.DB
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
	initRoutes(mux)
}

func main() {
	db, err := sql.Open("mysql", "go-n-blog:psw@tcp(localhost:3306)/goblog")
	if err != nil {
		log.Fatalln("DB connection failed. ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln("Ping to db error:", err)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Printf("Server online. Listening on %s\n", server.Addr)
	go server.ListenAndServe()
	for {
		//Menu
		fmt.Print(menuStr)
		s := bufio.NewScanner(os.Stdin)

		s.Scan()
		switch s.Text() {
		case "q":
			fmt.Printf("Closing server...\n")
			err := server.Shutdown(context.TODO()) //To study: what's a context? And how to use it properly?
			fmt.Printf("Server closed. %s", err)
			return
		case "t":
			fmt.Println("Realoading templates...")
			//Should we kill the server before?
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
