package main

import(
	"fmt"
)

func main(){

	// %s par indicar una cadena 
	nombre1 := "Juan"
	cadena1 := fmt.Sprintf("Hola, %s!", nombre1)

	fmt.Println(cadena1)


}


/* Return values
The fmt.Sprintf function returns just a single string that it has formatted.

Common specifiers
Following is a table of the most commonly used format specifiers in GO and their descriptions:

Specifiers	Description
%s	To print a string
%d	To print an integer
%v	To print values of all elements in a structure
%+v	To print the names and values of all elements in a structure

*/