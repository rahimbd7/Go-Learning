package main

import (
	"context"
	"db_crud/db"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		panic("Env file not found")
	}
	db.ConnectDB()
	defer db.Db.Close(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/users", getAllUsersHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("/users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Server is running on port: 5000")
	err = http.ListenAndServe("127.0.0.1:5000", mux)
	if err != nil {
		fmt.Println("Server Error", err)
	}

}
