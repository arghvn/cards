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
	Jim := person {
		firstname: "Jim",
		lastname: "party",
		contact: contactinfo {
			email: "Jim@gmail.com",
			zipcode: 94000,
		}
	}
	jim.print()
	// fmt.printf("%+v", Jim)
}

// output:
// firstname : jim lastname : party
// contact: {email:jim@gmail.com and zipcode : 94000}


// in deck application when we defined a new type of deck , some functions can be considered as a reciever
// like cards.shuffle()
// we can set same type by using struct(those have reciever)
// so in this code we can define a person as a reciever 
// we will define a function that recieve a person(like Jim)and do something like printing all persons details

func (p person) print() {
	fmt.printf("%v", p)
}

// above command means we can call this function with any type or any value related to person.

