package main

import "fmt"

func main() {
	var username string = "Chiraag Garg"
	fmt.Println(username)
	fmt.Printf("The username is of Type: %T \n", username)
	var tryempty string
	fmt.Println("Try Empty declaration is done here: ", tryempty)
	fmt.Printf("The Type of empty declartion is: %T", tryempty)
	var name string = "Chiraag Garg"
	var age int = 22
	fmt.Printf("Name: %s and Age: %d \n", name, age)
	//no_var_style
	users := 50
	fmt.Println("Number of Users:", users)
	fmt.Printf("Type of users is: %T", users)
}
