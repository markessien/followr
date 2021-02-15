package services

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/mmcdole/gofeed"
)

func UpdateFeeds(db *bolt.DB) {

	db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("Feeds"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {

			var feed Feed
			json.Unmarshal(v, &feed)

			fp := gofeed.NewParser()
			f, _ := fp.ParseURL(feed.Url)
			fmt.Println(f.Title)

			for i := 0; i < len(f.Items); i++ {
				fmt.Println(f.Items[i].Title)
				fmt.Println(f.Items[i])
			}
		}

		return nil
	})
}
