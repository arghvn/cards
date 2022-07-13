package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.exit(1)
	}
	fmt.println(resp)
}

// how to use interfaces in the standard library fo writting a program for this .
// we are going to write a program thats going to make an http requsest to google and then we will take just whatever
// response come back to us and we will print it out at the terminal

// http is a protocol that provide a space for get and send files like html file.

// GO will make the request it will then assign the response object to resp
// and if there was an error assigned to err.

// in error handling section we said that if something wrong happens we will exit the program
// exit is a command for closing the program
// overall if this code get 0 the program do its duty compeletly and close and if get any value except 0 this code
// close with error

type response struct {
	status     string // eg "ok"
	statuscode int    // eg 200
	proto      string // eg "http/1.0"
	protomajor int    // eg "1"
	protominor int    // eg "0"
}

// in GO we can take multiple interfaces
// different onterfaces and assemble thenm together to form another interface.
// both reader and closer are interfaces
// the reader read closer interface says if you want to fulfill the read closer interface like if you want to satisfy the requirements
// of this interface you have to satisfy the requirements of both are reader and the closer

// we defined two custom types englishbot and spanishbot with getgreeting function and return string

// type enflishbot struct
// func (englishbot)getgreeting() string

// we then difined an interface called bot and we said that if some other types inside of our application
// if any other types what soever defined a function called getgreeting that returned a strinng then that
// other type could also be considered to be a bot.

// it allowed us to then define a function called printgreeting
