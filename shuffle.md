shuffle function use for mix all the cards in the deck
this function take a deck and by order (a command) Arrange all cards randomly

for example we have Ace of spades , two of spades , three of spades  and four of spades
shuffle result : ' three of spades ' 'four of spades ' 'ace of spades ' 'two of spades'

in GO we dont have a standard library for shuffle the slice or array or another lists
then we do it manually

So we have to provide a number of cards and a number of indicators that fit the cards
and set each index belong to one card

for each index , card in cards
generate a random number between 0 and length of list (array or slice)-1
we want input all our cards in a loop and for each index slect a card

for example for ace of spades , two of spades , three of spades , four of spades ew need produce indexs 0 1 2 3 
each random number that generated , changed by first card in list
(if generate 3 , the first card belong to 3)
then the order of list : four of spades two of spades three of spades and ace of of spades
and 3 deletd becuase it used

swap the current card and the card at cards

for generate random number use math package and function rand in it
package rand implements random number generators
in this package we have intn function 
func intn(n int)int

