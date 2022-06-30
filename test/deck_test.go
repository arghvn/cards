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
}

// *testing.T is a test handler , verify the type of value that will send to the function

// out put : test passed