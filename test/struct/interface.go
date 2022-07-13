package main

import "fmt"

type bot interface {
	getgreeting() string
}

// type bot interface means every thing inside of this program all of you different types like english bot and spanish bot .
// all of you that may concern our program , our program now has a new type.

// we must define print related to greeting again
// this functions value (print Greeting function) call with b and related type is bot

// third type called bot that all languages can access to it.
type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	printGetgreeting(eb)
	printGetgreeting(sb)
}

// getGreeting() string :
// this command says if there is any another type inside of our program that has a function called getGreeting
// associated with it that returns a string then type is also returns a string then that type is also an honorary

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func printGetgreeting(sb spanishbot) {
	fmt.println(sb.getgreeting())
}

func (sb SpanishBot) getgreeting() string {
	return "Hola!"
}

// interface is a type

// output
// Hi there !
// Hola !

// We saw that we defined the printGreeting function once for two different sections.

// our program still has a type of english bot inside of it and it still has a typr of spanish bot
// and each of them has their own respective implementation of the greeting function.

// inside the deefinition of this type we said if you are a type in this program with a function called
// getGreeting and you return a type of string then you are now an  honorary member of type bot
// english bot and spanish bot either has a function that named getGreeting and both of them return a string


// now that we are also an honorary member of type bot , we can now call this function called printgreeting by 
  
func print Greeting (b bot){

}

// we said that any type inside of our program that tmplements a function calls getgreeting that returns a string
// and thats true of both english bot and spanish bot

// both of bots now included type bot 
// now any value of type english bot or spanish bot we can take those values and pass them to functions 
// of types bot as well.

// pritn Greeting says I only accept values of type bot 
// we passed in english bot and spanish bot to them

// we use interfaces to define a method set or a function set
// by defining this interface we are allowing GO to say hey mr developer any thing that kind of matches this description is now of type bot
// we can use those other types in any other location wherer we expect to see a type of bot

// in our program we have a type called english bot , engllish bot has a function called getgreeting associated with it.

// summary : 1-type english bot struct 2- func (englishbot getgreeting() string) 

// when we say associated we man to say that the getgreeting function expects to see a type english bot as its reciever 
// getgreeting is associated with the english bot type and then we did the exact same thing with spanish bot as well.

// we can use those other types in any other location where we expect to see a type of bot.  
// in our program we have a type called english bot...english bot has a function called getgreeting associated with it. 
// summary : 1-type english bot struct 2-func (englishbot getgreeting() string)

// when we say associated we can say that the getGreeting function expects to see a type englishbot as its reciever 
// getgreeting is associated with the english bot type and then we did the exact same thing with spanishbot as well ..

// summary : 1-type spanishbot struct 2-func(spanishbot)getgreeting() string  

// we then define a new interface called bot and we had said that this interface expects to see any other type inside of our applicaton that 
// implements a function called getgreeting and returns a string.  
// other type inside the application is then considered to also be of type bot.  
// type bot interface = getgreeting() string 
// both of english and spanish bot are bot type.  

// type bot interface {
// 	getgreeting (string,int)(string,error)
// }
// here ,bot = interfsce name and getGreeting is function name and string and int are list of argument types and string and error are list of return types

// we say type at first to tell GO that we then put the name of this new type ,where the name of the interface and then the interface keyword  
// inside the braces we list out all the different functions and the values that they are the types thst they have to return for another type to be considered  
// type bot.
// here be said the getgreeting function must accept an argument by type string and another argument by type int 

// Interfaces are unconditional (implicit) :
// we dont manually have to say that our custom type satisfies some interface 

// interface is a contract to help us manage types. 
// if our custom types implementation of a function is broken then interface wont help us.
// interface is a type that we can by its help ceate a special type  

// the word custom in interface means implicit 


type myinterface interfaace {
	my custom method()
}

func mycaller(my custom param my inteface){
	my custom params.my custom method()
} 

// in this code we have a function named mycaller
// that this function get a parameter named my custom params by type my interface 
// it can be by each type but has a condition
// it has to have a method by type my custom method

// so interfaces can read standard library documentations 