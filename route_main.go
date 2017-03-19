package main

import (
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
}

func index(writer http.ResponseWriter, request *http.Request) {
	data := NavBarData{
		Active: "home",
	}
	generateHTML(writer, data, "layout", "navbar", "home")
}

func about(writer http.ResponseWriter, request *http.Request) {
	data := NavBarData{
		Active: "about",
	}
	generateHTML(writer, data, "layout", "navbar", "about")
}
