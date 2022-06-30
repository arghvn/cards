
 run package tests run file tests
 run package test

 test in GO help us to check that our project run correctly or not.
 we want know how the testing works and make sure that result is that thing we excepted
 check for newdeck 
 about newdeck there are three points that if this points was ok means the testing is ok and the function works right
 1-each time a newdeck for each type of cards created , four state must be done (four string) ...for example spades , diamonds...
 (the length of deck must be 4 )
 2-the order must be define for example usually the first card should be Ace
 3-again the order must be define for example usually the last card should be Four
 for func test newdeck
 1-sode to make sure that a deck is created with x number of cards
 2-code to make sure that the first card as an Ace
 3- code to make sure that the last card is a Four of spades 

 for testing we can do it just for indicate in error template or show the solution as a notification
 write if statement to see if the deck has the right number of cards 
 if it doesnt , tell the GO test handler that something went wrong

suits {spades, diamonds, heart, club}
value {ace, two, three, four}
