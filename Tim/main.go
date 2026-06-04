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
// }t

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
// package main
// import "fmt"

// func sum(nums ...int) (s int) {
// 	for _, value := range nums {
// 		s += value
// 	}
// 	return

// }

// func main() {
// 	s := sum([]int{25,10,3,3,50,70}...)
// 	fmt.Println(s)
// }

///=====================================================
/// ADV GO

//Structs in Go
// usesd to group related data together
// field access trough dot noation only  user.name
// dynamic bracket notation does not exist without using complex reflection  user["name"]
//Anonymous structs are useful for temporary data structures

// In Go, you must declare a strict blueprint using the type and struct keywords before creating an instance. JavaScript allows you to just create objects inline on the fly.

// Blueprint is mandatory
// package main

// import "fmt"
// type User struct {
//     Name string
//     Age  int
// }

// func main() {
//     // Initialization
//     u := User{Name: "Alice", Age: 30}
// 	fmt.Println(u)
// }

/*
type StructName struct {
    fieldName1 type1
    fieldName2 type2
    fieldName3 type3
}*/

// exported name -> start with capital letter

// embedded structs
/*An embedded struct is created by declaring a field inside a struct without a field name (anonymous field). Go automatically pulls the inner struct's fields up to the outer struct. This is called field promotion.Syntax: Just the type name is listed.Behavior: You can access the inner fields directly as if they belonged to the outer struct.
JS Equivalent: It acts similarly to copying properties using the spread operator (...) or extending a class.*/

// package main

// import "fmt"

// type Dimensions struct {
//     Width  int
//     Height int
// }

// type Box struct {
//     Dimensions // Embedded struct (No field name given)
//     Material   string
// }

// func main() {
//     b := Box{Material: "Cardboard"}

//     // Field Promotion: Access width directly!
//     b.Width = 10

//     // You can still access it via the type name if you want:
//     b.Dimensions.Height = 20
// 	fmt.Println(b);
// }

//Interfaces:
// an interface is a named collection of method signatures that defines behavior without implementation.

// package main

// import "fmt"

// type Shape interface{
// 	getPerimeter() uint
// }

// type Triangle struct {
// 	sidea uint
// 	sideb uint
// 	sidec uint
// }

// func ( t Triangle) getPerimeter()  uint {
// 	return t.sidea + t.sideb + t.sidec
// }
// func (t Triangle) getSides() []uint{
// 	return []uint{t.sidea, t.sideb, t.sidec}
// }
// func main() {
// 	var s Shape = Triangle{1,2,3}
// 	fmt.Println(s.getPerimeter())
// }

// grabbing the perimeters of all of our difrent shapes

// implement all of the methods that are definned on it
// creating a flexible to be viewed
// Interfaces are all over the place and the main goal of them is that we abstarct away from the implementaion of our strucks

// Error Handling

//defer
// panic  - (won't run below code if we define panic)

// we can handle panic using the recover so its enable get rid from the panic error
// when recover occured , panic will not run

// and put recover in deferred function always

// ================================================

// Generics : Generics (also called type parameters) were added in Go 1.18. They allow you to write functions and data structures that work with multiple types while maintaining type safety.

// package main

// import "fmt"

// // Generic function - T is a type parameter
// func PrintSlice[T any](s []T) {
//     for i, v := range s {
//         fmt.Printf("Index %d: %v\n", i, v)
//     }
// }

// func main() {
//     // Works with any type
//     nums := []int{1, 2, 3, 4}
//     strs := []string{"a", "b", "c"}

//     PrintSlice(nums)      // Type inference: T = int
//     PrintSlice(strs)      // Type inference: T = string
// }

/*
	func FunctionName[TypeParam TypeConstraint](params) returnType {
	    // Use TypeParam as a type
	}
*/
// package main

// import "fmt"

// type GenericSlice[T any] []T
// func (g GenericSlice[T]) Print()  {
// 	for _, value := range g{
// 		fmt.Println(value)
// 	}
// }

// func main() {
// 	g:=GenericSlice[int]{1,2,3}
// 	g.Print()
// }
