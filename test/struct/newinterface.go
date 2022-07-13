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
