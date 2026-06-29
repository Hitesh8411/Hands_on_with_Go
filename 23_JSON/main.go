package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"Name"`
	Age int 	`json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main() {
	fmt.Println("Json")
	fmt.Println("we are learning JSON")
	person := Person{Name: "John", Age:34,IsAdult: true}
	fmt.Println("Person data is :",person)

	//convert person into JSOn Encoding (marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshling", err)
		return
	}
	fmt.Println("Person data is:", string(jsonData))


	//Decoding (Unmarshalling)
	var personData Person
    err = json.Unmarshal(jsonData,&personData)
	if err != nil {
		fmt.Println("Error unmarshiillaing",err)

	}
	fmt.Println("person detail is after unmarshalling :", personData)
}