// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }


// type conversions

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x, y int = 3.0, 4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)
// 	fmt.Println(x, y, z)
// }

// type interface

// package main

// import "fmt"

// func main() {
// 	v := 42 // change me!
// 	fmt.Printf("v is of type %T\n", v)
// }

// constants

// constants cannot be decaled hby suing the := syntax.

// package main

// import "fmt"

// const Pi = 3.14

// func main() {
// 	const World = "世界"
// 	fmt.Println("Hello", World)
// 	fmt.Println("Happy", Pi, "Day")

// 	const Truth = true
// 	fmt.Println("Go rules?", Truth)
// }

// Numreic constants 
// package main

// import "fmt"

// const (
// 	// Create a huge number by shifting a 1 bit left 100 places.
// 	// In other words, the binary number that is 1 followed by 100 zeroes.
// 	Big = 1 << 100
// 	// Shift it right again 99 places, so we end up with 1<<1, or 2.
// 	Small = Big >> 99
// )

// func needInt(x int) int { return x*10 + 1 }
// func needFloat(x float64) float64 {
// 	return x * 0.1
// }

// func main() {
// 	fmt.Println(needInt(Small))
// 	fmt.Println(needFloat(Small))
// 	fmt.Println(needFloat(Big))
// }

// for 
// init and post statements are optional

// if 
// same as like for loops
// not surrounded by parentheses()
// but the braces {} are required

// loops and functions
// package main

// import (
// 	"fmt"
// )

// func Sqrt(x float64) float64 {
// 	z := 1.0
	
// 	for 1=0; i < 10; i++{
// 		z -= (z *z -x) / (2*z)
// 	}
// 	  return z
		
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Print("Go runs on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("macOS.")
// 	case "linux":
// 		fmt.Println("Linux.")
// 	default:
// 		// freebsd, openbsd,
// 		// plan9, windows...
// 		fmt.Printf("%s.\n", os)
// 	}
// }

// switch evaluation order 
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("When's Wednesday?")
// 	today := time.Now().Weekday()
// 	switch time.Wednesday{
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// // }

// switch with no condition
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening.")
// 	}
// }

//defer  : 
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

// package main

// import "fmt"

// func main() {
// 	defer fmt.Println("world")

// 	fmt.Println("hello")
// }

// package main

// import "fmt"

// func main() {
//     defer fmt.Println("Step 3: I run last!") // Pushed to defer stack
    
//     fmt.Println("Step 1: I run first.")
//     fmt.Println("Step 2: I run second.")
// }

// stacking defers 

// deferredfunctions calls are pushed onto a stack .when a function returns ,its defered calls are executed in last-in-first-out order 

// package main

// import "fmt"

// func main() {
// 	fmt.Println("counting")

// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}

// 	fmt.Println("done")
// }

// Structs is a collection of fields.
// Strut fields are accessesd using a dot 

// package main

// import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	v.X = 4
// 	fmt.Println(v.X)
// }

// Pointer to structs 

// struct fields can be accessed through a struct pointer 

// struct literals

// is an expression that crates and intializes a new struct instances directly in you code 


//Array
// slices : doesnt store any data 
// changing the elements of a slice modifies the corresponding elements of its underlying array.

// slices are like refrencing to arrays
package main

import "fmt"

func main() {
	names := [4]string{
		"John",//0
		"Paul",//1
		"George",//2
		"Ringo",//3
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// Slice literals
// To declare a slice literal, look at the empty square brackets []. Go calculates the size and handles the memory allocation automatically.

// package main

// import "fmt"

// func main() {
// 	q := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(q)

// 	r := []bool{true, false, true, true, false, true}
// 	fmt.Println(r)

// 	s := []struct {
// 		i int
// 		b bool
// 	}{
// 		{2, true},
// 		{3, false},
// 		{5, true},
// 		{7, true},
// 		{11, false},
// 		{13, true},
// 	}
// 	fmt.Println(s)
// }

// slice default :

// we can set the low bound as zerp and the length of the underlying slice or array for the high bound

// a[0:10]
// a[:10]
// a[0:]
// a[:]


// 3,5,7
// 3,5
// 2


//0
//2,3,5,7