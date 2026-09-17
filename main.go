package main

import "fmt"

type address struct {
	city    string
	country string
}

type user struct {
	name string
	age  int
	address address
}
/*Receiver functions*/
	func (u user) greet() {
		fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.name, u.age)
	}
// func add(a int, b int) int {
// 	return a + b
// }
// func twoNumbers(a int, b int) (int, int) {
// 	sum := a + b
// 	mul := a * b
// 	return sum, mul
// }

// func stringPrint(name string) {
// 	fmt.Printf("Hello, %s!\n", name)
// }

// //multiple return values
// func multipleReturnValues(age int, city string) (int, string) {
// 	return age, city
// }
// func add(x *int) {
// 		*x=11
// 		fmt.Println("The value of x after increment is", *x)
// 	}
func main() {
	// fmt.Println("Hello, World!")
	// a := 10
	// b := 20
	// sum := a + b
	// fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)

	// logged := false
	// if !logged {
	// 	fmt.Println("User is logged in")
	// }
	// age := 0
	// if age > 20 {
	// 	fmt.Println("You are an adult")
	// } else if age == 20 {
	// 	fmt.Println("You are exactly 20 years old")
	// } else {
	// 	fmt.Println("You are a minor")
	// }

	//swtich case
	// a := 3

	// switch a {
	// case 1:
	// 	fmt.Println("a value is 1")
	// case 2:
	// 	fmt.Println("a value is 2")
	// default:
	// 	fmt.Println("a value is not 1 or 2")
	// }

	// fmt.Println(add(10, 20))
	// sum, mul := twoNumbers(10, 20)
	// fmt.Printf("The sum is %d and the multiplication is %d\n", sum, mul)

	// stringPrint("Rahim!")
	//multiple variables
	// x,y:= 10, 20
	// fmt.Println("The value of x is", x, "and the value of y is", y)
	// multiple return values
	// age, city := multipleReturnValues(25, "New York")
	// fmt.Printf("Age: %d, City: %s\n", age, city)

	//annonymous function and IIFE (Immediately Invoked Function Expression)
	// func() {
	// 	fmt.Println("This is an anonymous function")
	// }()
	

	//for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("The value of i is", i)
	// }


	//user input
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	// fmt.Printf("Hello, %s!\n", name)

	//array
	// var employees [3]string
	// for i := 0; i < len(employees); i++ {
	// 	fmt.Printf("Enter the name of employee %d: ", i+1)
	// 	fmt.Scanln(&employees[i])
	// }
	// for i := 0; i < len(employees); i++ {
	// 	fmt.Printf("Employee %d: %s\n", i+1, employees[i])
	// }

	//short hand array declaration
	// employees := [3]string{"Alice", "Bob", "Charlie"}
	// for i := 0; i < len(employees); i++ {
	// 	fmt.Printf("Employee %d: %s\n", i+1, employees[i])
	// }
	//array slice
	// ids := []int{1, 2, 3, 4, 5}
	// newIds := ids[1:4]
	// fmt.Println("Original IDs:", ids)
	// fmt.Println("Sliced IDs:", newIds)

	//pointers in  function

	// add:=func(x *int) {
	// 	*x=11
	// 	fmt.Println("The value of x after increment is", *x)
	// }
	// y := 10
	// add(&y)
	// fmt.Println("The value of y is", y)

	//array pointer in function parameter
	// modifyArray:= func (arr *[3]int) {
	// 	arr[0] = 10
	// }
	// myArray := [3]int{1, 2, 3}
	// fmt.Println("Before modification:", myArray)
	// modifyArray(&myArray)
	// fmt.Println("After modification:", myArray)

	//structs
	// user1 := user{name: "Alice", age: 30}
	// fmt.Printf("User1: %+v\n", user1)
	// fmt.Println("User1 Name:", user1.name)
	// fmt.Println("User1 Age:", user1.age)


	//embedded structs
	// personDetails := user{name: "John", age: 30, address: address{city: "New York", country: "USA"}}
	// fmt.Printf("Person Details: %+v\n", personDetails)
	// fmt.Println("Person Name:", personDetails.name)

	/*constructor function */
	// createUser := func(name string, age int, city string, country string) user {
	// 	return user{name: name, age: age, address: address{city: city, country: country}}
	// }
	// user2 := createUser("Bob", 25, "Los Angeles", "USA")
	// fmt.Printf("User2: %+v\n", user2)
	// user2.greet()

	/*Map */
	// userMap := make(map[string]user)
	// userMap["user1"] = user{name: "Alice", age: 30, address: address{city: "New York", country: "USA"}}
	// userMap["user2"] = user{name: "Bob", age: 25, address: address{city: "Los Angeles", country: "USA"}}

	// userMap2:= map[string]user{
	// 	"user1": {name: "Alice", age: 30, address: address{city: "New York", country: "USA"}},
	// 	"user2": {name: "Bob", age: 25, address: address{city: "Los Angeles", country: "USA"}},
	// }
	// fmt.Printf("User Map 2: %+v\n", userMap2)


	/*range in map,slice,string,array, channel*/
	// myMap := map[string]string{
	// 	"name":    "Alice",
	// 	"country": "USA",
	// 	"city":    "New York",
	// }
	// for key, value := range myMap {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }
	// fmt.Println("------------")
	
	// for _, value := range myMap {
	// 	fmt.Println(value)
	// }
	// //range in slice
	// colors := []string{"Red", "Green", "Blue"}
	// for index, color := range colors {
	// 	fmt.Printf("Color %d: %s\n", index, color)
	// }
	// //range in string
	// str := "Hello"
	// for _, char := range str {
	// 	fmt.Printf("%c \n", char)
	// }

	

}
