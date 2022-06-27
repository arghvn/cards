package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type og 'deck'
//which is a slice of strings

type deck []string

//With the above command code, we declared that wherever a string is created in the slices,
// it is ready and they can be replaced with deck type.
// In fact, we told the compiler that we have something called deck type that is equal to slice of  string behavior

//slice of string in main.go file 
//cards := []string{"Ace of Diamonds",Newcard()}

deck = []string

//We want to create a new function that belongs to the deck, and whenever we try to print a new card,
// we consider the process as a loop in the main file.
//We want to use receivers, for examlpe  d deck

func (d deck) Print(){ 

	for i , cards := range d{

		fmt.Println(i,cards)

	}

}

//Since both files are connected, go to the main file and delete its loop, which is written according to the following command, 
//and write the following command instead.

//for i ,cards := range cards {

	//fmt.Println(i,cards)

//}

//cards.Print()

//Note that inside the body of the main function, we defined a variable called cards, 
//which we then put in print. Here, with the changes we made, we transferred the cards variable as a d variable to the print function. 
//According to the contract, one or two letter abbreviations are assigned to the recipients, so it can be written:
//func (cards deck)

d = cards 


//Access to print any variable of type deck :
//func (d deck) print() {}

//Newdeck  : This function returns a list of new cards for us, It is also a deck type

func Newdeck () deck {

}

//Question : Does the function we defined above need a receiver ?
//Response : No The purpose of this function is to create a Newdeck variable.
// If we wanted to work on the deck we needed to.
// So receivers are only used where we want to call a variable from the existing variables.

cards := deck {

}

//With the above command, we created a variable called cards and of the deck type.
// This variable starts with zero and initially there are no cards in it.
//This zero value will eventually reach 52 cards
//So we have to have the cards mixed in them
//for example :

cards := deck {"Ace of spades ","Two of spades "}

//This seems like a long and time consuming task
//As a clever method, you can create an empty deck and then create two separate slices in them
//One of these slices holds the cards(suit) and the other slice stores the value of each(value).

//So we can set two different loops to repeat in the suit list and the card value list.
// And determine each time a card is combined with a value of a value and a value of a suit and go to the carddeck.
// First we need to create a slice of the strings and insert the types of cards in it.

//The deck is a temporary place where playing cards go in it.

cardsuits := []string{"Spades","hearts","Diamonds","club"}
cardsvalue := [string{"Ace","Two","Three","four"}

//We need to write two for loops nested. 
//The final state is a variable called a card that has a suit attribute with a value attribute.

//Finally, we want to return the slice card, then use the return command.


for i, suit := range cardsuits {
	for j, value := range cardValues {
		cards = append(cards, value+" of "+suit)
	}
}


//In this case, changes i and j are not used.
//The compiler gives us an error with the message that
//i and j declared and not used
//To fix the error, we have to tell the compiler that we are aware of these bugs,
// but we do not intend to use these variables at the moment, 
//so we use the  underline.

//The number that displays :4*4=16

//deal function : This function designs a new hand for the game
//create a new 'hand' of cards

//deck : This function stored all the playing cards in an array
//with deal function we create a new deck 

//deal function : The idea of this function is that we take an existing batch of playing cards 
//and then create a new batch of a certain size.

//for example : We have a deck of 4 cards, Ace of spades, Two of spades, Three of spades, Four of spades
//Now we determine that we want to have a folder (hand) of cards containing three cards
//deal(3)
//Our hand consists of three of the four available cards and one of the cards remains inside the deck.
//The remaining card type is still the deck
//deal : So with this function, we divided a slice into several slides.

//save to file function : The purpose of this function is to save a deck to the hard disk of the machine used for programming.

// in byte slice we have here , there are amont of string type 
// in past we have a deck and change it to slice of string []string
// then change it to string and finally change to []byte (byte slice)
// now take this steps opposite
// we have []byte changing to string and changing to []string and then change to deck
// []byte slice show all character in game
// changing []slice to string 
string(bs)
// by this command we can see byte slice as string

// if in thisb step we print the values we have Ace of spades , two of spades ...
// joining all of this do with strings package 
// but for oppisite we also use strings package 
// in this package we have a function named split

// func split(s, sep string) []srtring
// this function change a slice of string to multple seperated strings

s := Strings.split(string(bs), ",")
// s is a slice of string
// in fact we assigned slice of string to va variable named s
// at first we write the type we want and then the type we have
return deck(s)

func shuffle (d deck) {
	for i := range d {
		newPosition := rand.Intn(len(d)- 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
// we can use instead of d , the length of array or slice that each number that we want for us
