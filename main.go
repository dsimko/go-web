package main

import (
	"net/http"
	"log"
	"time"
)

type NavBarData struct {
	Active string
}

func main() {

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("resources"))
	mux.Handle("/resources/", http.StripPrefix("/resources/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/about", about)

	log.Println("Listening...")

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}
