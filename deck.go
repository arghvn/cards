package main

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
