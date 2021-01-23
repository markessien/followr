package main

import (
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
	if name := r.FormValue("site"); name != "" {

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			err := b.Put([]byte("answer"), []byte("42"))
			return err
		})

	}
}
