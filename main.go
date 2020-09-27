package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type message struct {
	MsgType string
	MsgText string
}

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tmpl, err := template.ParseFiles("templates/index.html")
			if err != nil {
				fmt.Println(err)
			}
			tmpl.Execute(w, nil)
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm error: %s", err)
			}
			//login := r.FormValue("login")
			log.Println("post form")
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
