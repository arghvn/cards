package main

import "testing"

func TestNewDeck(t *testing.T) {

	// this function use for testing NewDeck
	// automatically call by GO test runner , this parameter refer to an argument that we refer to it by T
	d. := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
		// each newdeck has 4 suit and 4 value
		// in the end must has 16 cards generated
		// in this commands we said if the function length not equal to 16 a parameter went worng
		// we notice this note to testhandler (responsible for testing)
		// the text used for this reason
		// the second argument show this authorizes number for client
	}
	// testing two other parameters has a same code :
	if d[0] != "ace of spades"{
		t.Errorf("expected first card of ace of spades, but got %v", d[0])
	}
    if d[len(d) - 1] != "four of clubs"{
		t.Errorf("expected last card of four of clubs, but got %v", d[len(d) - 1])
	}
}

// *testing.T is a test handler , verify the type of value that will send to the function

// out put : test passed


func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// this function use for testing eveything related to save and delete.
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveTofile("_decktesting")
	loadedDeck := newDeckfromfile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("expectes 16 cards", in deck, got %v, len(loadedDeck))
	}
	os.Remove("_deckktesting")
}

// output : pass

// for making sure that the test work correctly , we make a changes in it for examlpe enter 160 instead 16

// output :
// error
