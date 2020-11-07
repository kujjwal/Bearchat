package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterRoutes(router *mux.Router) error {
	router.HandleFunc("/api/profile/{uuid}", getProfile).Methods(http.MethodGet)
	router.HandleFunc("/api/profile/{uuid}", updateProfile).Methods(http.MethodPut, http.MethodPost)

	return nil
}

func getUUID(w http.ResponseWriter, r *http.Request) (uuid string) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		http.Error(w, errors.New("error obtaining cookie: " + err.Error()).Error(), http.StatusBadRequest)
		log.Print(err.Error())
		return
	}
	//validate the cookie
	claims, err := ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, errors.New("error validating token: " + err.Error()).Error(), http.StatusUnauthorized)
		log.Print(err.Error())
		return
	}
	log.Println(claims)

	return claims["UserID"].(string)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	// Obtain the uuid from the url path and store it in a `uuid` variable
	// Hint: mux.Vars()
	// YOUR CODE HERE
	UUID := mux.Vars(r)["uuid"]

	// Initialize a new Profile variable
	//YOUR CODE HERE
	profile := Profile{}

	// Obtain all the information associated with the requested uuid
	// Scan the information into the profile structs' variables
	// Remember to pass in the address!
	err := DB.QueryRow("SELECT * FROM users WHERE uuid=? LIMIT 1", UUID).
		Scan(&profile.Firstname, &profile.Lastname, &profile.Email, &profile.UUID)
	
	/*  Check for errors with querying the database
		Return an Internal Server Error if such an error occurs
	*/
	if err != nil {
		http.Error(w, errors.New("error querying profile database").Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

  	//encode fetched data as json and serve to client
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		http.Error(w, errors.New("error encoding profile").Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	return
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	
	// Obtain the requested uuid from the url path and store it in a `uuid` variable
	// YOUR CODE HERE
	UUID := mux.Vars(r)["uuid"]

	// Obtain the userID from the cookie
	// YOUR CODE HERE
	userID := getUUID(w, r)

	// If the two ID's don't match, return a StatusUnauthorized
	// YOUR CODE HERE
	if userID != UUID {
		http.Error(w, errors.New("incorrect access token").Error(), http.StatusUnauthorized)
		log.Println(errors.New("incorrect access token").Error())
		return
	}

	// Decode the Request Body's JSON data into a profile variable
	profile := Profile{}
	err := json.NewDecoder(r.Body).Decode(&profile)

	// Return an InternalServerError if there is an error decoding the request body
	// YOUR CODE HERE
	if err != nil {
		http.Error(w, errors.New("error decoding profile").Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	// Insert the profile data into the users table
	// Check db-server/initdb.sql for the scheme
	// Make sure to use REPLACE INTO (as covered in the SQL homework)
	_, err = DB.Exec("REPLACE INTO users VALUES (?, ?, ?, ?)",
		&profile.Firstname, &profile.Lastname, &profile.Email, &profile.UUID)

	// Return an internal server error if any errors occur when querying the database.
	// YOUR CODE HERE
	if err != nil {
		http.Error(w, errors.New("error updating profile in DB").Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	return
}
