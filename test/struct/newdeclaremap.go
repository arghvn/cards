// definition map with variables

package main

import "fmt"

func main () {
	var colors map[string]string{
	// map above means the assigned type is indicated
	fmt.Println(colors)
	}
}

// output :
// map[] is empty map , it means no pair key value dosent exist in it
// because whenever we define a variable in GO and dont assign value to it , GO automatically indicate 
// initial value as 0

// the third way is direct way 

// consider printmap function and assign a argument to it (c) ant then define this argument to colormap

type map
type key string
type value string

func printmap(c map[string]string){
	for color,hex := range c {
    fmt.println("hex code for", color ,hex ,color ,"is" ,hex)
	}
	printmap(colors)
}

// in this function we use commands for iteration over map by use for key and then use two variables for recieve each key and value in loop
// colors = key , hex = value

// so to define GO that we want iterate over map and the for body length is constant ,we can use commands that be unique for key and value pair.

// summary :
// c is an argument ,color is key for this step through the loop ,hex is value for this step through the loop


// output :
// hex code for green is #45f745
// hex code for white is #ffffff
// hex code for red is #ff0000

// difference between struct and map
// in map all keys must be the same type but in structs all the values can be of completly different types.
// in structs values can be of different types.
// this comparsion just do for value because in structs , there are is no meaning named key.

// in maps keys are indexed , so we can iterate over then ,so we can iterate over all pairs by a map.
// this note is impossible in structs

// important note :
// map is reference type and struct is value type

// when we assign a map to a function in fact we are assigning a reference to a datastructure. 
// for example printmap function that the thing we have was colors , if we create a change in printmap like * or & or make a change in datastructure 
// in main section of colors : we will see the updated output

// so to assign map to function doesn use copy but also use reference
// when we assign a struct to a function , in fact we create a copy from all of the struct
// when this struct changed the original section not changed because this section define out of function. 


