// new way to define structs

package main

type contactinfo struct {
	email   string
	zipcode int
	contact contactinfo
}
type person struct {
	firstname string
	lastname  string
}

func main() {
	alex.firstname = "Aelex"
	alex.lastname = "Anderson"
}

// %+v this command assign all alex properties to it as value

// we define a variable by type person and assign alex value to it
// sofar nothing feature or property not assigned to alex

// if in GO we define a variable and dont assign property to it ,GO assign it 0 automatically
// for example for string , "" that means empty string
// and if it was int or float ,assign 0
// for bool ,false
// this values name zero value
// in output we see zero values as {} for string

//for this code
// func main () {
// 	var alex person
// 	fmt.println("%+v", alex)
// }
// output :
// {firstname : lastname}
// two strings with empty value just have name

// for this code
// func main () {
// 	alex.firstname = "Aelex"
//    alex.lastname = "Anderson"
// }
// output :
// {firstname : Alex lastname : Anderson}

// other features in structs:
// one of the best properties is embeding , embeding a struct inside another struct

// for example in our struct each person need to connection way , we will define contact
// firstname in struct was string
// lastname in struct was string
// contact in struct is contactinfo
// contactinfo properties : email string and zipcode int
// we difine second struct and inside the main struct make a contact and call second struct on it.

// person definition in struct3
