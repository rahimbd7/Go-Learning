package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var currentUsers = []User{
	{
		Id:    1,
		Name:  "rahim",
		Email: "rahim@mail.com",
		Age:   34,
	},
	{
		Id:    2,
		Name:  "kahim",
		Email: "kahim@mail.com",
		Age:   24,
	},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/users", userHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("/users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Server is running on port: 5000")
	err := http.ListenAndServe("127.0.0.1:5000", mux)
	if err != nil {
		fmt.Println("Server Error", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to HTTP CRUD")
}

// create a user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	newUser.Id = len(currentUsers) + 1
	currentUsers = append(currentUsers, newUser)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

// get a single user
func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid id")
		return
	}

	for _, user := range currentUsers {
		if user.Id == id {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(user)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}

// update a single user
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid id")
		return
	}
	var updateUser User
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid body request")
		return
	}

	for idx, user := range currentUsers {
		if user.Id == id {
			updateUser.Id = id
			currentUsers[idx] = updateUser
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}

// delete a user
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid id")
		return
	}
	for idx, user := range currentUsers {
		if user.Id == id {
			currentUsers = slices.Delete(currentUsers, idx, idx+1)
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(currentUsers)
	// users,_:=json.Marshal(currentUsers)
	// w.Write(users)
}
