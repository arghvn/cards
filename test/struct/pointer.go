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

	func (pointerToperson *person) updatename(newfirstname string) {
		(*pointerToperson).firstname = newfirstnameconst
	}
	jim.updatename("Jimmy")
	jim.print()
}

func (p person) updatename(newfirstname string) {
	p.firstname = newfirstname
}



func (p person) print() {
	fmt.printf("%v", p)
}


// output:
// firstname : Jimmy ,lastname : party ,contactinfo = {email:Jim@gmail.com and zipcode : 94000}
// so the upadate function work corrcetly