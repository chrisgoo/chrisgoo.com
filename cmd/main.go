package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"text/template"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("error: ", err)
	}
}

func run() error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/template/index.html"))
		tmpl.Execute(w, nil)
	})
	fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe("127.0.0.1:8080", nil)

	return nil
}
