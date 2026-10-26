package main 

import "fmt"

func divide(a,b float64) (float64,error) {
	if b == 0 {
		return 0, fmt.Errorf("denomintor must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("started Error handling in the Functions")
    
	ans,err:= divide(10,5)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Println("Divison of two numbers is",ans)
}

//Go rule : if you decalared any variable, package in go you have to use 