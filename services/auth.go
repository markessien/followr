package services

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/boltdb/bolt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Password     []byte
	EmailAddress string
	SessionToken string
	SessionTime  int
	LoggedIn     bool
	FeedIDs      []int
}

func CreateUser(db *bolt.DB, emailAddress string, password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	db.Update(func(tx *bolt.Tx) error {

		var user User
		user.EmailAddress = emailAddress
		user.Password = hashedPassword

		user_json, err := json.Marshal(user)

		b := tx.Bucket([]byte("Users"))
		err = b.Put([]byte(emailAddress), user_json)
		return err
	})
}

func LoginUser(db *bolt.DB, emailAddress string, password string) (User, error) {

	var user User

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		v := b.Get([]byte(emailAddress))

		err := json.Unmarshal(v, &user)

		// Compare the stored h66ashed password, with the hashed version of the password that was received
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			// If the two passwords don't match, return a 401 status
			return nil
		}

		// Create a new random session token
		sessionToken := uuid.NewV4().String()
		user.SessionToken = sessionToken

		// Set the token in the cache, along with the user whom it represents
		// The token has an expiry time of 120 seconds

		b = tx.Bucket([]byte("Sessions"))
		err = b.Put([]byte(sessionToken), []byte(user.EmailAddress))
		if err != nil {
			// If there is an error in setting the cache, return an error
			return errors.New("Could not save session")
		}

		return nil
	})

	return user, nil
}

func ValidateLoggedIn(db *bolt.DB, w http.ResponseWriter, r *http.Request) (User, error) {

	var user User
	user.LoggedIn = false

	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return user, errors.New("No Cookie")
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return user, errors.New("Bad Request")
	}
	sessionToken := c.Value

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Sessions"))
		v := b.Get([]byte(sessionToken))

		email_address := string(v)

		b = tx.Bucket([]byte("Users"))
		v = b.Get([]byte(email_address))

		err = json.Unmarshal(v, &user)
		user.LoggedIn = true
		return nil
	})

	return user, nil
}

func RefreshSession(w http.ResponseWriter, r *http.Request) {

	/*
		// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		response, err := cache.Do("GET", sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if response == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// (END) The code uptil this point is the same as the first part of the `Welcome` route

		// Now, create a new session token for the current user
		newSessionToken := uuid.NewV4().String()
		_, err = cache.Do("SETEX", newSessionToken, "120", fmt.Sprintf("%s", response))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Delete the older session token
		_, err = cache.Do("DEL", sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `session_token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   newSessionToken,
			Expires: time.Now().Add(120 * time.Second),
		})

	*/
}
