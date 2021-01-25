package main

import "net/http"

func handle_routes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/add-site", add_site)
	http.HandleFunc("/login", login)
}
