package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func serveStatic(writer http.ResponseWriter, request *http.Request) {
	log.Println("Endpoint Hit: serveStatic")

	indexTemplate, err := template.ParseFiles("web/index.html")
	if err != nil {
		fmt.Println(err)
	}
	_ = indexTemplate.Execute(writer, nil)
}
