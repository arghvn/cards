so in GO slices are wat we refer to as a reference type 
reference type = slices , maps , channels , pointers , functions

value type in GO = int , float . string . bool . struct  ...  use pointers to change these things in a function
it means if we have some value int , flot or struct ... and get it to a function , when GO see reference types , take a copy from it and store main value in memory
so when we use these we dont worry about using pointers.

introduction new feature in GO , map
its very similar to structs
in Go map is a collection of key value pairs...like dictionary un python
dictionary in python is a data type (a set of elements without index and order)
map has a pair of key and value related together

when we import many key , they must be in one type ... each key related to a value
as well all values that related to key must be in one type
but key and value not essential be in same type

in key section and value section we have different two set of types 
