// package main 

// import "fmt"

// func addFunction(a, b int) int {
// 	return a + b
// }

// func main() {
// 	fmt.Println("function in Golang")
// 	add := addFunction(3,4)
// 	fmt.Println("add of two number is", add)

// }

// package main

// import "fmt"

// func greet(name string) {
// 	fmt.Println("Hello,", name)
// }

// func add(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	greet("Amit")
// 	sum := add(10, 20)
// 	fmt.Println("Sum:", sum)
// }

package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}