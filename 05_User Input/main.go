package main

import "fmt"

func main() {
	fmt.Println("Learning about how to import packages and take input!")

	// 1. Declare a variable to store the input
	var name string

	fmt.Print("Enter your first name: ")
	
	// 2. Use &name to pass the memory address so Scanln can modify it
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)
}

// using bufio 
// package main
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Learning about how to import packages and bufio!")

// 	// 1. Create a reader that listens to Standard Input (the keyboard)
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your favorite quote: ")

// 	// 2. Read until the user hits the 'Enter' key (\n)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	// 3. Clean up the trailing newline character
// 	input = strings.TrimSpace(input)

// 	fmt.Printf("Your quote is: \"%s\"\n", input)
// }

//When using fmt.Scanln(&name), that little & symbol is a pointer. It tells Go, "Hey, don't just look at the name variable; look at the exact spot in memory where name lives and write the user's input directly into it."