package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FantasyApi/interfaces/users"
)

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Endpoint Hit: create")
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data to the server")
	}

	// if empty body {}
	var user users.Users
	json.NewDecoder(r.Body).Decode(&user)
	if user.IsEmpty() {
		json.NewEncoder(w).Encode("No data found")
		return
	}
	// Append to the user table
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
	// Send a 200 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("user Created")

}
