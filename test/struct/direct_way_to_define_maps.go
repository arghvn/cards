package main

import "fmt"

func main() {
	colors := make(map[string]string){
    fmt.println(colors)
	}
}

// we want use a internal function in GO to create map that say create map
// this command create a new map for us that exist no value in it  

// if here we have a defined map and we want assign values to it , at first write the name of amp and then key and then value

colors["white"] = "ffffff"

// output :
// map[white : #ffffff]

// a important function in map issue is delete function
// it use for deleting value an key
// at first we can value and then secod argument(key)

delete(colors, 10)

// output :
// map[]

// static variable
// It means that the value attributed to that variable should be set only once, that is, only one location of the memory is assigned 
// to such variables. In other parts of the program, the said variable cannot be set again, but the execution of a series of
// commands related to this variable leads to The change in value is attributed