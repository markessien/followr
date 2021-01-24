package services

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

type Site struct {
	Name     string
	Url      string
	Password string
}

func AddNewSite(db *bolt.DB, site_name string, site_password string) {
	if site_name != "" && site_password != "" {

		db.Update(func(tx *bolt.Tx) error {

			var site Site
			site.Url = site_name
			site.Password = site_password

			site_json, err := json.Marshal(site)

			b := tx.Bucket([]byte("Websites"))
			err = b.Put([]byte(site_name), site_json)
			return err
		})

	}
}
