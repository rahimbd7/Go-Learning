package main

import (
	"context"
	"db_crud/db"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

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

	query := `
	insert into users(name,age,email)
	values ($1,$2,$3)
	returning id
	`
	if err = db.Db.QueryRow(context.Background(), query, newUser.Name, newUser.Age, newUser.Email).Scan(&newUser.Id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Couldn't Create user...")
		return
	}

	// currentUsers = append(currentUsers, newUser)
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

	query := `
	select id,name,age,email from users where id=$1	
`
	var user User
	if err := db.Db.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
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
	query := `
	update users 
	set name=$1, age=$2, email=$3
	where id= $4
	returning id,name,age,email
	`
	if err = db.Db.QueryRow(context.Background(), query, updateUser.Name, updateUser.Age, updateUser.Email, id).Scan(&updateUser.Id, &updateUser.Name, &updateUser.Age, &updateUser.Email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Couldn't Found the user", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateUser)
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
	var user User
	query := `delete from users where id=$1`
	if err := db.Db.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
}
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	query := `select * from users`
	rows, err := db.Db.Query(context.Background(), query)

	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintln(w, "could not find the users")
		http.Error(w, "Couldn't found the users", http.StatusInternalServerError)
		return
	}
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email); err != nil {
			http.Error(w, "failed to retrieve user", http.StatusInternalServerError)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Couldn't found the users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)

	// users,_:=json.Marshal(currentUsers)
	// w.Write(users)
}
