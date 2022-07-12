every value that we declare in GO has to have a type assigned to it 
even if for each value we can say each value has a type

in GO each function that be defined , it must exist with argument and type of argument and its reciever
every function has a type
every value has a type 
every function has to specify the type of its argument or reciever 

for interface introduction we consider a program that this program create two seperated chat robot , the first use foor English and 
the second use for Spanish.
this two robots have same functions and same structs
we are going to imagine that both of these bots have a function that takes its recpective type as a reciever and the functions name is getgreeting

chatbot1 = type englishbot struct
chatbot2 = type spanishbot struct

chatbot1 = func(englishbot1getgreeting()string)
chatbot2 = func(spanishbot1getgreeting()string)

in both of this bots ,the function getgreeting function return for us a string
for example chatbot1 = hello there and chatbot2 = hola (hello in spanish)

both bots are also going to have a function tied to them called print greeting

chatbot1 = funcprintgreeting(eb englishbot)
chatbot2 = funcprintgreeting(sb spanishbot)

the only difference between this two robot is the function argument

eb is short for english bot
sb is short for spanish bot

the main code in getgreeting functions will be different

getgreeting in englishbot include codes that specially create a greeting in english.
so these robots have different function

keep on in chatbot.go file
