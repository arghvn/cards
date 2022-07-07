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