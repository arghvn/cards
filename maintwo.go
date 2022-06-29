package main

func main() {
	cards := newdeck()
	caeds.shuffle()
	cards.print()
}

// output in terminal :
// created a list of random cards and show it

// each time we mix the deck of card , usually GO provided by this list unchanged means in same state but randomly

rand : Intn(len(d) -1)
// intn is random number generator
// GO by default uses a way to generate random number
// This method depends on parameters such as seed and its value.
// seed is a source for generate random number
// we import seed to generator numbers then generator create a set of random (numbers or any thing else)

// The problem here is that GO always uses a specific and similar seed by default.
// So instead of changing the whole program, we should look to change the values in the seed.
// Each time the program is run, a new seed is created and this seed is used to generate numbers.
// refer to GO pacakges / package rand and type rand section
// in fact type rand is a source for random number (an object for generate random numbers)
// we want it generates a set of random number for us and assign it to seed.

// func New(src source) *rand
// new returns a new rand that uses random values frommm src to generate other random values.

// source in this package : 
// a source represents a source of uniformly distributed pseudo random int 64 values in the range.

// this package tell us for generate a source we must call a newsource and a set of number from type int64
// first define an object named source (seed)
// then take a commnd that update this seed randomly

// by this command seed generating depend on for loop 
// r generate  a value by type rand

// in this package we have func(r *rand) Intn(n int) int
// it means if we  have a value by type rand , this value can call int function and we put this value in int and then return.
// we have r by type rand and assign it to Newposition (all must be int64)

// each time we run this code , this values are different 
// for connecting values to time of running code we use time package
// func Now() time 
// Now returns the current local time 
// another releted function named unix nano
// func (t time) unix nano() int64
