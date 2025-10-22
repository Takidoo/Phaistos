package Server

import (
	"html/template"
	"net/http"
)

// Route Handles

func GameHandle(w http.ResponseWriter, r *http.Request) {
	return
}

func AdminHandle(w http.ResponseWriter, r *http.Request) {
	return
}

func ProfileHandle(w http.ResponseWriter, r *http.Request) {
	return
}

func Testhandle(w http.ResponseWriter, r *http.Request) {
	var tmpl, _ = template.ParseFiles("test.html")
	tmpl.Execute(w, nil)
}
