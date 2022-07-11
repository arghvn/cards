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

now we want to know after creating slice , how the defined values save in RAM
when in GO we define a slice ,GO automatically create a slice data struture for us that has 
length ,capacity and pointer(pointer to the array) in a memory address

myslice = []string{Hi...} , address = 0001 and slice {length,cap,point to haed} , address = 0002 and point to head = []string{Hi...}
in fact myslice dont refer to 0002
this array named underlying array in GO
so when we refer to myslice , GO returns the slice in 0001 (not array in 0002)

now if we call a function and assign myslice to it , GO still work as a pass by value language 
it means GO take a copy srom that value and save it
so when we call update slice and assign it to our slice , in fact we take slice data structure and copy it and 
save in another memory address

even though the slice data structure is copied it is still pointing at the original array in memory
because the slice data structure and the array data structure are two seperated elemnt s in memory.

we copy slice but it point to arayy still 

so when we modify this array(0002) or when we inside the function when we attempt to modify the slice we sre 
still modifying the same array that both copies of the slice are now pointing to 

in GO there are several other types of elements that behave in exactly the same way where the underlying data structure or the value the
value that we create is kind of this what we refer to as a reference type .
the slice data structure is what we refer to as a reference type because it is a refernce to another data structure in memory

its totally ok if we make a copy of this refernce in memory because its still always going to be pointing back to the same underlying true
source of datao

