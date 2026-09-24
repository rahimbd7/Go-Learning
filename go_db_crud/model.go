package main

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
