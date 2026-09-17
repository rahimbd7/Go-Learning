package main
import "fmt"

//any = interface{} //any type of data
//type assertion = converting any type of data to a specific type
//ok to check if the conversion is successful or not

func processData(data any) {
     userData,ok:= data.(int) //type assertion
     if ok {
         fmt.Println("Data is of type int:", userData)
     }
     userName,ok:= data.(string) //type assertion
     if ok {
         fmt.Println("Data is of type string:", userName)
     }else {
         fmt.Println("Data is not of type string")
     }
}
//init function

//unlimited number of parameters
func add(numbers ...int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

//greets to unlimited number of users
func greetUsers(greetTitle string,users ...string) {
    title := greetTitle
    for _, user := range users {
        fmt.Println(title, user)
    }
}
func main() {  
    // processData(10)
    // processData("hello")
    // processData(false) //not of type string

    // result := add(1, 2, 3, 4, 5)
    // fmt.Println("Sum:", result)

    // greetUsers("Alice", "Bob", "Charlie")
    users:= []string{"David", "Eve", "Frank"}
    greetUsers("Hello, Welcome!",users...) //passing slice as variadic argument
}