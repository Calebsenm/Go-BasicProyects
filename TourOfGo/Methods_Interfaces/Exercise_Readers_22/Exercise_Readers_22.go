/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (mr MyReader) Read ( b []byte ) ( int , error ){
	i := 0
	p := func () int{
		b[i] = 'A';
		i += 1
		return i 
	}
	
	fmt.Println(p)
	return p() , nil

}

func main(){

	reader.Validate(MyReader{})

}