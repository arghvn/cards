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