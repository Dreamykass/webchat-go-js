package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	Port = ":8080"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		fmt.Println(err)
	}
	_ = t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", serveStatic)
	_ = http.ListenAndServe(Port, nil)
}
