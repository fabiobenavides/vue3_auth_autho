package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Register struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

var users []*User

func registerUser(w http.ResponseWriter, r *http.Request) {
	var register Register

	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := 1
	for _, user := range users {
		if user.ID >= id {
			id = user.ID + 1
		}
	}

	user := &User{
		ID:       id,
		FullName: register.FullName,
		UserName: register.UserName,
		Password: register.Password,
	}

	users = append(users, user)

	saveUsers()

	respondJSON(w, http.StatusOK, nil)
}

func loadUsers() (err error) {
	content, err := ioutil.ReadFile("users.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(content, &users)
}

func saveUsers() (err error) {
	content, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("users.json", content, 0o644)
}
