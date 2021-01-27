package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/markessien/followr/services"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

func index(w http.ResponseWriter, r *http.Request) {

	// templates := template.Must(template.ParseFiles("templates/index.html"))
	templates, _ := template.New("").ParseFiles("templates/index.html", "templates/base.html")

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}

	if err := templates.ExecuteTemplate(w, "base", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func site_timeline(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.ParseFiles("templates/site_timeline.html"))
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}

	if err := templates.ExecuteTemplate(w, "site_timeline.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	site_name := r.FormValue("sitename")
	site_password := r.FormValue("sitepassword")

	services.AddNewSite(db, site_name, site_password)

	templates, _ := template.New("").ParseFiles("templates/signup.html", "templates/base.html")
	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {

	templates, _ := template.New("").ParseFiles("templates/login.html", "templates/base.html")
	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func add_feed(w http.ResponseWriter, r *http.Request) {
	site_name := r.FormValue("sitename")
	feed_url := r.FormValue("feed_url")

	services.AddNewSite(db, site_name, feed_url)

	templates := template.Must(template.ParseFiles("templates/add.html"))
	if err := templates.ExecuteTemplate(w, "add.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
