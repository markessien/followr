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

	user, _ := services.ValidateLoggedIn(db, w, r)

	m := map[string]interface{}{
		"User":  user,
		"Feeds": nil,
	}

	// templates := template.Must(template.ParseFiles("templates/index.html"))
	templates, _ := template.New("").ParseFiles("templates/index.html", "templates/base.html")

	if err := templates.ExecuteTemplate(w, "base", m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	// user := nil

	user, err := services.ValidateLoggedIn(db, w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	m := map[string]interface{}{
		"User":  user,
		"Feeds": nil,
	}

	// templates := template.Must(template.ParseFiles("templates/index.html"))
	templates, _ := template.New("").ParseFiles("templates/dashboard.html", "templates/base.html")

	if err := templates.ExecuteTemplate(w, "base", m); err != nil {
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
	user_email := r.FormValue("email")
	user_password := r.FormValue("password")

	if user_email != "" && user_password != "" {
		services.CreateUser(db, user_email, user_password)
	}

	templates, _ := template.ParseFiles("templates/signup.html", "templates/base.html")
	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {

	user_email := r.FormValue("email")
	user_password := r.FormValue("password")

	if user_email != "" && user_password != "" {
		user, err := services.LoginUser(db, user_email, user_password)

		if err == nil {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   user.SessionToken,
				Expires: time.Now().Add(1200 * time.Second),
			})

			http.Redirect(w, r, "/dashboard", 301)
		}
	}

	templates, _ := template.New("").ParseFiles("templates/login.html", "templates/base.html")
	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func add_feed(w http.ResponseWriter, r *http.Request) {
	user, err := services.ValidateLoggedIn(db, w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// site_name := r.FormValue("sitename")
	feed_url := r.FormValue("feed_url")

	services.AddNewFeed(db, feed_url, user.EmailAddress)

	http.Redirect(w, r, "/dashboard", 301)
}
