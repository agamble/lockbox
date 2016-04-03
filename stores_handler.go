package main

import (
	"html/template"
	"net/http"
)

func StoresHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("templates/stores.html")
		t.Execute(w, nil)
	case http.MethodPost:
	default:
	}
}
