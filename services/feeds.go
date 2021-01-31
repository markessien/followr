package services

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

type Feed struct {
	ID      int
	Name    string
	Url     string
	AddedBy string
}

func AddNewFeed(db *bolt.DB, feed_url string, user_email string) {
	if feed_url != "" {

		db.Update(func(tx *bolt.Tx) error {

			var feed Feed
			feed.Url = feed_url
			feed.AddedBy = user_email

			b := tx.Bucket([]byte("Feeds"))

			v := b.Get([]byte(feed_url))
			if v == nil {
				id, _ := b.NextSequence()
				feed.ID = int(id)

				feed_json, _ := json.Marshal(feed)
				b.Put([]byte(feed_url), feed_json)
			}

			b = tx.Bucket([]byte("Users"))
			v = b.Get([]byte(user_email))

			var user User
			json.Unmarshal(v, &user)

			found := false
			for i := 0; i < len(user.FeedIDs); i++ {
				fmt.Printf("%s\n", user.FeedIDs[i])

				if user.FeedIDs[i] == feed.ID {
					found = true
					break
				}
			}

			if found == false {
				user.FeedIDs = append(user.FeedIDs, feed.ID)
				user_json, _ := json.Marshal(user)
				b.Put([]byte(user_email), user_json)
			}

			return nil
		})

	}
}
