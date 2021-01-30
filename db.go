package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func init_db() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = bolt.Open("db/follows.db", 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	db.Update(func(tx *bolt.Tx) error {

		// A single new item found in a feed
		_, err = tx.CreateBucketIfNotExists([]byte("Users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("Sessions"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err := tx.CreateBucketIfNotExists([]byte("Websites"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("Feeds"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("WebsiteFeed"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		// A single new item found in a feed
		_, err = tx.CreateBucketIfNotExists([]byte("FeedItem"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil
	})
}

func GetValueFromDB(db *bolt.DB, table string, key string) {

}
