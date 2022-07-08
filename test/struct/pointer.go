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
	

	Jimpointer := &Jim
	Jimpointer.updatename("Jimmy")
	Jim.print()

	}
	// in Jimpointer := &Jim we use & and assign Jim to Jimpointer and call Jimpointer
	
	Jim.updatename("Jimmy")
	Jim.Print()


	func (pointerToperson *person) updatename(newfirstname string) {
		(*pointerToperson).firstname = newfirstnameconst
	}
	jim.updatename("Jimmy")
	jim.print()
}

func (p person) updatename(newfirstname string) {
	p.firstname = newfirstname
}

// this function can call with any reciever by type pointToperson (exactly whatever exist in Jimpointer)


func (p person) print() {
	fmt.printf("%v", p)
}


// output:
// firstname : Jimmy ,lastname : party ,contactinfo = {email:Jim@gmail.com and zipcode : 94000}
// so the upadate function work corrcetly

// we call Jimpointer.updatename (memory address transfer to this as pointerToperson)
// Jimpointer = pointer to a person

// (*PointerToperson) = memory address , it is a actual pointer
// if says take this memory address pointer the person and turn it into an actual value.

// Jimpointer(an address like 0001) := Jim(value like person {firstname : "JIm ..."}) , change value to pointer
// turning address into value with address is possible 
// turning value into address with value

// for change pointer to value do (*pointerTOperson.firstname....)


// whenever we see the star(*) at a location where normally we would specify a type that means we are looking for a type that is a pointer to this (to a person) thats
// this pointer to a person variable should be .

// so in this case the star should not really be thought of as an operator
// its only when we have an actual memory address or an actual pointer to write (*pointerToperson) does that the star operator turns that back into a value.

// when sign * come before a type (like type person ,it means we are looking for a type of pointer that refer to a variable that store a pointer
// when the sign * come before actual pointer or a variable that store a pointer ,it means we want change this actual pointer to a person value.




// bot now we will notice that there is a mismatch in our reciever types.
// in fact vs code tell us Jim is a variable by type person and Jim has a value from struct
// and there is not any address memory and pointer here 
// but we can use updatename function that has a reciever by type pointerToperson