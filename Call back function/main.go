package main

import (
	"encoding/json"
	"fmt"
)

//defer value stay the unchange even the value updated in code after defer call But in closure defer function copy the reference of variables and defer function always call in delay

func calculator(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func test(x int) {
	fmt.Println("value in defer function", x)
}

// normal return
func deferInClosure() int {
	result := 10
	defer func() {
		result += 100
		fmt.Println("inside defer closure:", result)
	}()

	fmt.Println("value of result", result)
	result += 100
	return result

}

// name return where its store result  address
func nameReturnDeferClosure() (result int) {
	result = 10
	defer func() {
		result = result + 100
		fmt.Println("inside defer closure:", result)
	}()
	fmt.Println("value of result", result)
	result = result + 100
	return
}

//Enum

type Weekend int

const (
	Monday Weekend = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func officeDay(day Weekend) {
	switch day {
	case Sunday, Monday, Tuesday, Wednesday:
		fmt.Println("Office Open")
	case Thursday:
		fmt.Println("Halfday")
	case Friday, Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("invalid day")

	}
}

//string enum

type userRole string

const (
	ADMIN     userRole = "ADMIN"
	USER      userRole = "USER"
	MODERATOR userRole = "MODERATOR"
)

func checkUserRole(role userRole) {
	fmt.Println("Your Access role:", role)
}

type user struct {
	Name string `json:"personName"` //struct tag
	Age  int	`json:"-"`
	City string `json:"city"`
}

func main() {
	// add := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(calculator(10, 20, add))
	// fmt.Println(calculator(20, 20, add))

	// value := 10
	// defer test(value)
	// fmt.Println("updated value",value+10)

	// finalResult := deferInClosure()
	// fmt.Println("Final return from closure value",finalResult)
	// fmt.Println("================================================")
	// finalResult2:=nameReturnDeferClosure()
	// fmt.Println("Final return from closure value",finalResult2)

	// officeDay(Thursday)
	// checkUserRole(USER)

	//struct to json

	u := user{
		Name: "rahim",
		Age:  23,
		City: "ctg",
	}
	rawJson, err := json.Marshal(u)
	if err!=nil{
		fmt.Println("Errors:",err)
	}
	// fmt.Println(string(rawJson))

	//json to struct
	var u2 user
	err= json.Unmarshal([]byte(rawJson),&u2)
	if err !=nil{
		fmt.Println("errors")
	}
	fmt.Println(u2)
}
