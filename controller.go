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

	templates := template.Must(template.ParseFiles("templates/index.html"))
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}

	if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
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

func add_site(w http.ResponseWriter, r *http.Request) {
	site_name := r.FormValue("sitename")
	site_password := r.FormValue("sitepassword")

	services.AddNewSite(db, site_name, site_password)

	templates := template.Must(template.ParseFiles("templates/add.html"))
	if err := templates.ExecuteTemplate(w, "add.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
