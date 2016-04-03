package main

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("%+v", res)
	log.Printf("%+v", req)
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, nil)
}
