package main

// in robot we must create struct and then getgreeting function and printgreeting function

type EnglishBot struct{}
type SpanishBot struct{}

// struct{} is struct without any associated fields.

// in struct we have to define the ability to recieve Englishbot and Spanishnot as a reciever

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	printGetgreeting(eb)
	printGetgreeting(sb)
}

func printGetgreeting(eb englishbot) {
	fmt.println(eb.getgreeting())
}
func printGetgreeting(sb spanishbot) {
	fmt.println(sb.getgreeting())
}

func (eb EnglishBot) getgreeting() string {
	// very custom logic for generating an English greeting
	return "Hi there!"
}

func (sb SpanishBot) getgreeting() string {
	return "Hola!"
}

// in both function we see the reciever section that recieve eb or sb .

// function getgreeting for english and spanish probably very different logic in these functions

// function printgreeting ,this will probably have identical logic
