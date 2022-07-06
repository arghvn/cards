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

	jim.updatename("Jimmy")
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

// we want to define a function that get person as a reciever and change the firstname 
// we should in this function argument recieve firstname as a string

func (p person) updatename(newfirstname string) {
	p.firstname = newfirstname
}

// this function has a reciever by type person and replacement is done successfully

func (p person) print() {
	fmt.printf("%v", p)
}

// above command means we can call this function with any type or any value related to person.

// in terminal we have output :

// firsname: jim
// lastname and contact info without changes 

// a summary about this commit :
// definition a function that recieve a struct as a reciever
// definition print and update for name that both of them have a reciever by type person

// the name not changed 
// this commit goal : why does after definition functions like Update and name the task dont work correctly ?
// pointer 
// in computer RAM is like a set of boxes , each boxes in our computer can have data  and value ...
// each data and value in this boxes has a seperated and unique address.
// sometime like this code the compilre tell the computer i want retrieve some of data , this retrievation done in RAM
// RAM look at to addresses and that address use for storing data and take down value from it

// in programm we define a type by name type person and assign it to a person by name Jim
// Jim := person{}
// after do this , GO in that struct create a person for that person


