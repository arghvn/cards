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
	fmt.printf("%+v", Jim)
}

// output:
// firstname : jim lastname : party
// contact: {email:jim@gmail.com and zipcode : 94000}
