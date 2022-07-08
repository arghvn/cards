& and * are operators

&variable : ampersand operator 
&Jim = Jimpointer
Jim is pointing at our struct sitting in memory and that struct is existing at some particular RAM address
for example the address memory is 0001 , if we call , assign the value to pointer
this created pointer doesnt refer to struct directly  but also to the address contain struct
so if we print Jimpointer , the value is shown 0001

*pointer : star operator 
*memory address = pointer
take this memory address and give me the value exit at that memory address
in this case pointerToperson is the memory address that Jim exit at

*pointerToperson : we are saying heres the pointer (jimpointer = 0001)

I dont want to look at the memory address any more.
instead give me direct access to whatever this thing or whatever value is actually string here

*Jimpointer turns into the actual struct of type person.
Jim is a refrence to the struct in memory (the actual vaalue of the struct)

we used & its cause Jim changed to memory address or a pointer and then assign it as a value to Jimpointer
then we called Jimpointer.update and then print it
in function definition we changed reciever type to *person (a pointer that points at a person)

when we use * and a pointer , pointer change to value

*person : a type of pointer that point at a person

note : 
func(pointerToperson *person(this is a type description it means we are working with a pointer to a person))updatename(){
    *pointerToperson(this is an operator it means we want to manipulate the value the pointer is referencing)
}


(pointerToperson *person) in the reciever we said that pointer to a person is a value of type *person

*person , it means that is a type description and it means that this update name function can
 only be called with the reciever of a pointer to a person
this is a specifing that we want a type of type pointer to a person

*pointerToperson : this is an actual operator that takes this pointer and returns it into an actual value.
