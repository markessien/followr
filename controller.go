package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
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

func add_site(w http.ResponseWriter, r *http.Request) {
	if site_name := r.FormValue("site"); site_name != "" {

		db.Update(func(tx *bolt.Tx) error {
			fmt.Println("Updating")
			b := tx.Bucket([]byte("Websites"))
			err := b.Put([]byte("site-name"), []byte(site_name))
			return err
		})

	}

	templates := template.Must(template.ParseFiles("templates/add.html"))
	if err := templates.ExecuteTemplate(w, "add.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
