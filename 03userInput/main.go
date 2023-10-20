package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	username := "chiraag"
	fmt.Println("User Name: ", username)
	reader := bufio.NewReader(os.Stdin)
	println("Enter the age of the user:")
	age, _ := reader.ReadString('\n')
	fmt.Println("Age:", age)
}
