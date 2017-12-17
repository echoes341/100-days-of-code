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
	var err error
	db, err = sql.Open("mysql", "go-n-blog:psw@tcp(localhost:3306)/goblog")
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

func userTestInsert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := user{
		ID:       0,
		Name:     "Test Name",
		User:     "test@test.te",
		Password: "encryptedpassword,indeed",
	}
	err := u.insert()
	if err != nil {
		fmt.Fprintf(w, "Coulnd't insert the user: %s\n", err)
	} else {
		fmt.Fprintf(w, "User added!\n")
	}
}

func printUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("SELECT * FROM users;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Select users error: %s", err)
	}
	var u user
	for rows.Next() {
		rows.Scan(&u.ID, &u.Name, &u.User, &u.Password)
		fmt.Fprintf(w, "ID: %d   NAME: %s   USERNAME: %s   password: %s\n", u.ID, u.Name, u.User, u.Password)
	}
}
