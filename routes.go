package main

import "net/http"

func handle_routes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/add-feed", add_feed)
	http.HandleFunc("/dashboard", dashboard)
}
