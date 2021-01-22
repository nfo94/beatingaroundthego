package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

// CreateUser inserts a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read request body"))
			return
    }

    var user user

    if err = json.Unmarshal(){

		}
}
