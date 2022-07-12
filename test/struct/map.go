package main

import "fmt"

// goal : creating a map and find how its work
// we want define a map named colors
// key type in it is tyhpe string and value as well
// we define map and assign map to it
// define a map that all keys in it are string (in[]) and all values in it is string (out of[])

func main () {
	colors := map[string]string{
		// we want assign a string as a key named red to it and in front of it , write the value
		"red": "#ff0000"
		"green": "#4bf745"
	}
	fmt.println(colors)
}

// in fact we write this map in order to relate each color to a hexcode
// each color has a hexcode , hexcode for red in google = ff000
// hex means hexadecimal
// That is, numbers in base 16
// In this naming system, 16 symbols are used, including numbers and letters
// The application of this system is to color the program

// output in terminal :
// map[red:#ff0000 green:#4bf745]

// so we defined a map and assign initial value to it
// another way to define map
// newdeclaremap.go file

// where we can use a datastructure against another ? means where we can use struct for datastructure and wher map ?
// when we want indicate a set of related property , we use map 
// for example :
// this code is mapping between each color and related hex

// in map all values must be the same type .
// in struct keys dont support indexing .
// map use to represent a collection of related properties 
// struct use to represent a thing with a lot different properties . 
// in map we dont need to know the list  of all different keys or all different fiels in complie time. 
// in struct we need to know all the different fields at compile time.

// in map we can create a map and add or cut some properties.
// in struct the properties must be define correctly at first.
