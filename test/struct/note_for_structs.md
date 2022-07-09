In the structures, when we gave the desired changes to GO, a copy of the structure was created inside GO.
Therefore, when we change that function, we change the copy and not the original structure.

the difference between slice and array in GO :

slice can grow annd shrink 
slice used 99% of the time for lists of elements 

arrays primitive data structure cant be resized
arrays rarly used directly

GO is directly access to slice and array

in less time we use array because there is initial data in it and dont have resize property
mostly dont occur that we want create a list of things with constant length
We often need a list of data that can change according to our code
so we consider slice as a fancy array

so when we create a slice ,GO create two seperated data structures for us.
for example for this slice : myslice := []string{"Hi there..."} ,GO create two data structures for us 
the first is slice 
each slice has three component {pointer ,capacity ,length}
pointer = is how many elements it can contain at present
capacity = in fact this slice pointer is a pointer to array that show the really list of items.
length = is represents how many elements currently exist inside the slice

array = Hi there , how are you?


