package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

// we define person type and assign values

func main() {
	Alex := person{firstname: "Alex", lastname: "Anderson"}
	fmt.Print(Alex)
	// fmt.Print(Alex) just use for debug error(declare and not use)
}

// tell GO what fields the person(in our struct) has first name string , last name as well
// create a new value of type person
// for examlpe firstname is Alex and lastname is Anderson

// in main function we identify people
// for declare Alex we must use :=
// because we are initializing and assigning value

// output :
// Alex Anderson
