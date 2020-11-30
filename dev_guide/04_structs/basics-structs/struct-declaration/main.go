package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// declaring empty struct
	//var alex person

	// better way
	alex := person{firstName: "Alex", lastName: "Anderson"}

	// printing information contained in struct
	fmt.Printf("%+v", alex)
}
