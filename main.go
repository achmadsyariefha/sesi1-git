package main

import (
	"fmt"
)

func main() {
	// name, address := "John Doe", "123 Main St"
	// age := 30
	// fmt.Printf("Hello, %s\n", name)
	// fmt.Printf("Age: %d\n", age)
	// fmt.Printf("Address: %s\n", address)
	// fmt.Printf("Type %T, %T, %T\n", name, age, address)

	name := "John Doe"
	age := 30
	isMale := true
	fmt.Printf("name: %s\nage: %d\nisMale: %t\n", name, age, isMale)
	fmt.Printf("\nData Type: \n%T\n%T\n%T\n", name, age, isMale)
}