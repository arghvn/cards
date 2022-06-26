package main

import (
	"fmt"
    "os"
)
func main() {

	// newdeck, This function creates a new batch 
	// of cards each time it is called


	//An array is required in which each element 
	//represents a card from the game

    //For each function we define, we must define a task

    //We need a function that can print cards, we name it print

	//The function we use for the mixing property is called shuffle

	//To select a set of cards, we use the deal function

    //We select a variable of type string and assign a value to it
	//The name of this variable is card


	cards := []string{"Ace of Diamonds",Newcard()}

	//Easy way to define variables 
	//cards :="Ace of spades"
	//colon equal

	cards.Print()

	//We want to define a function that assigns different strings
	// to a variable because we have 52 cards

	func Newcard () string {

		
    cards := Newcard ()

	//In line 6, we change the variable from one mode to 
	//several modes with this command card := Newcard
	
	//About the definition of the Newcard function :
	//Go asks us to state very clearly what kind of data to return from each function
	//We told the compiler to consider a certain string type whenever a new card is called.

	return "five of diamonds"

	}

	//If we did not write the last command, return : While saving the file, it shows us an error that you 
	//have written a function that does not return any data
	//error: No result value excepted  or  error: too many arguments to reurn, have (string) and want()
	//In low-level languages, it must be clear what the function returns
	//The second error says there are too many arguments to return while the string is empty.

	//Analyze this command: func Newcard () string {} :
	//define a function by name Newcard
	//Assign value to it as a string
	//In the return section, we also want to return the string
	//If we wanted to return the int variable, for example, we would have to write reutrn 12 without " "

	//This command only creates one card while we need several cards so we go to the arrays and slices
	//for example:

	Primes := [6] int {2,3,5,7,11,13}

	var s[] int = Primes[1:4]


	fmt.Println(s)

	//Slice S refers to the values of the array prime

	t := int {2,3,5,7,11,13}

	//The length of the array is always greater than or equal to the slice

	//the append function, With this function we can add an item to the slice,
	// the length of the array is always constant so replacement occurs.

	//In the array we defined, only 5 cards could be stored, while if we defined a slice,
	// we could increase or decrease them by mixing.

	//So we have to make a slice that has several cards in it
	//So instead of the commandو var card.string = "Ace of spades" و In the body of the main function
	//The variable must be of type string and replace this command in the body of the main functionو cards := []string{}
	

	//For example, so that our slice has two new cardsو cards := []string{Newcard(),Newcard()}

	//cards := []string{"Ace of Diamonds",Newcard()}, The final file we wrote is to add a card for us

	//A common arrogance in development: cards declared and not used 

	//We have to print it to fix the error و fmt.Println(cards)

	//With the append function, we can add a member (card) to the slice

	cards = append(cards,"saix of spades")

	//If we type the print command, we have three outputs , six of spades , cards(ace of diamonds and newcard)

	//loop for
	//Temporarily delete the print command in the commands (fmt.Print(cards)
	//and add the loop in the body of the main function
	//we have three outputs, (0)Ace of diamonds, (1)Five of Dimanods, (2)six of Diamnods
	//indexes start from zero

	//We need to create a new type of data type here
	//We want to expand a base type and add some extra features to it
	//decktype creates that corresponds to strings
	//This type of data keeps the list of playing cards in the form of a slice for us, each member of which is a string
	//type deck []string
	//The above command tells the compiler that we want to create an array of strings 
	//and add a set of functions that are used to work with them.

	//To display the code that shows how the deck works, create a go file in the project folder
	
	//In the first line, the connection between the folders in this project must be specified with the pacakemain command

	//In fact, we created this extra type of data so that we can create custom functions with it that can only be executed with this type of data.
	// This type of data is very similar to slices.

    fmt.Println("Error:", err)
	
	// use os package for exit 
	// this standard library peovides a platform independent interface to operating system functionality
	// it means no different what is our operating system , 
	// use exit from this package and assign 0 argument to it 
	// assigning 0 means the code run correctly and assigning each value except 0 means we have a problem

	os.Exit(1)






	



