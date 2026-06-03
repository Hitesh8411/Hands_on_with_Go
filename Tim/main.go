// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	x := "11110111"
// 	// y,err := strconv.Atoi(x)
// 	y,err := strconv.ParseInt(x,2,0)
// 	fmt.Println(y,err)
// }

//ParseInt , ParseBloat, ParseBool

// switch case

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 1

// 	switch a {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	default:
// 		fmt.Println("default")

// 	}
// }

// in go we can exclude the variable

// fallthrough : automatically pass doen to the next case

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := -2

// 	switch {
// 	case a < -1:
// 		fmt.Println("a is less than -1")
// 		fallthrough
// 	case a < 0:
// 		fmt.Println("a is less than 0")
// 		fallthrough
// 	case a < 1:
// 		fmt.Println("a is less than 1")
// 	default :
// 	    fmt.Println("deafult case")

// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := "b"

// 	switch {
// 	case a == "a", a == "b", a == "c":
// 		fmt.Println("a is a, b or c")
// 	default :
// 	    fmt.Println("deafult case")
// 	}
// }

// go doest have while loop

// package main

// import "fmt"

// func main() {
//     i := 0
//     for i < 5 {  // This acts as a while loop
//         fmt.Println(i)
//         i++
//     }
// }

// _ annonyms placeholder
//looping over the string -> for range

// Array :
// package main

// import "fmt"

// func main() {
// 	var arr [2]bool
// 	fmt.Println(arr)
// }

// func main(){
// 	arr := [2]int{1,2}
// 	fmt.Println(arr)
// }

// func main() {
// 	arr := [2][2]int{{1,2},{2,3}}
// 	fmt.Println(arr)
// }

// func main() {
// 	arr := [...][2]int{{1,2},{2,3}}  // [...]its exlipict account the counting of number of arrays
// 	// its doesn't mean size is dynamic [...] calculate the size of array auto.

// 	// mutate the array
//     arr[0] = [2]int{10,11}  // create a new array to mutate

// 	fmt.Println(arr)
// 	fmt.Println(len(arr))

// 	for i, value := range arr {
// 		fmt.Println(i, value)
// 	}
// }
// all of the values inside of my array need to be the exact same type.
/*
we can't change the size of array
its is fixed sizze which means when we intiates
it is always going to be that size
we can add, delate and change whats inside of it

*/

// func main(){
// 	arr := [...][2]int{{1,2},{2,3},{3,2}}

// 	for _, nested := range arr {
// 		for _, value := range nested {
// 			fmt.Println(value)
// 		}
// 	}

// }

// func main() {
// 	arr := [...][2]int{{1,2},{2,3},{3,2}}
// 	test(arr)
// 	fmt.Println(arr)
// }

// func test(arr [3][2]int) {
// 	arr[0] = [2]int{100, 100}
// }

// here we are not able to modify the underlying array
//opt [[1 2] [2 3] [3 2]]

// also the size of the array is associted with the type of that array.

// Slice : a slice is a view of an array.

// package main
// import "fmt"

// func main() {
// 	arr := [5]int{1,2,3,4,5}
// 	sl := arr[:3]    // 1 2 3
// 	sl1 := arr[1:3]  // 2 3
//     // sl3 := arr[-1:5]
// 	sl[0] = 100
// 	fmt.Println(sl, arr)
// 	fmt.Println(sl1)
// 	// fmt.Println(sl3)
// 	fmt.Println(sl1[0]) // this work on updated view

// }

// func main() {

// 	arr := [...]int{11,22,32,34,65,66,67,68}

// 	slice := arr[4:]

// 	fmt.Println(arr)
// 	fmt.Println(slice)
// }

// func main() {
//      //pointer -> arr[0]  , arr[1]
// 	 //length ->  2
// 	 //capacity  ->  4

// 	arr := [5]int{13,14,15,16,17}
// 	// sl := arr[:3]
// 	   sl := arr[:3]
// 	   sl = sl[:4]
// 	fmt.Println(sl, "len(sl)\n", cap(sl))
// }

// func main() {
// 	s1 := []string{"hello","wprld"}
// 	// []string -> slice type
// 	//[2]string -> array type
// 	fmt.Printf("%T", s1)
// }

// func main() {
// 	sl := []string{"hello","world"}

// 	for x := 0; x < 10; x ++ {
// 		sl = append(sl,"Hitesh")
// 		fmt.Println(sl, len(sl), cap(sl))
// 	}

// 	fmt.Println(sl)
// }

// some more ways to create a slice function

// func main() {
// 	sl := make([]int,10,20)
// 	fmt.Println(sl)
// }

// Iterate over a slice

// func main() {
// 	sl := []string{"hello","world","hi"}

// 	for i, value := range sl {
// 		fmt.Println(i, value)
// 	}

// }

//  func main() {
// 	s1 := []string{"hitesh", "jayesh", "pritesh"}

// 	fmt.Println(s1)
// }

// func test(arr []string){
// 	arr[0]="suryvananshi"
// }

// if your passas array you're not going to mutate it ,
// whereas if you pass a slice , you can mutate that

// Map

// package main

// import "fmt"

// func main() {
// 	// var mp map[string]int = map[string]int{"a":1} 
//      mp := map[string]int{"a":1}

// 	fmt.Println(mp)
// }


// creating a MAP

// package main
// import "fmt"

// func main() {
// 	mp := make(map[string]int)
// 	fmt.Println(mp)
// }

// func main() {
// 	mp := map[string][]int{"a":{1,2,3}}
// 	fmt.Println(mp)
// }

// package main

// import "fmt"

// func main() {
//     // Create map
//     countryCapitalMap := make(map[string]string)
    
//     // Add key-value pairs
//     countryCapitalMap["France"] = "Paris"
//     countryCapitalMap["Italy"] = "Rome"
//     countryCapitalMap["Spain"] = "Madrid"
    
//     fmt.Println(countryCapitalMap)
//     // Output: map[France:Paris Italy:Rome Spain:Madrid]
    
//     // Update existing key
//     countryCapitalMap["France"] = "Lyon"  // Overwrites
//     fmt.Println(countryCapitalMap["France"])  // Output: Lyon
// }


///Functions
// have to define the types of our Functions:


// package main 
// import "fmt"

// func add(num1 int, num2 int) int{
// 	return num1 + num2
// }

// func main() {
// 	value := add(1,2)
// 	fmt.Println(value)
// }


// --
// package main 
// import "fmt"

// func callFunc(callable func(int) int) int {
// 	return callable(10)
// }

// func doubleNumber(num int) int {
// 	return 2 * num 
// }

// func main() {
// 	value := callFunc(doubleNumber)
// 	fmt.Println(value)
// }


// 

// package main 
// import "fmt"

// func sum(nums ...int) int {

// }

// func main() {
// 	s := sum(4,5,6)
// 	fmt.Println(s);
// }


// named returnd values
package main
import "fmt"

func sum(nums ...int) (s int) {
	for _, value := range nums {
		s += value
	}
	return

}

func main() {
	s := sum([]int{25,10,3,3,50,70}...)
	fmt.Println(s)
}

