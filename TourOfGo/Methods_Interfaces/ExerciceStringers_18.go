package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr ) String() string{
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main(){
	hosts := map[string]IPAddr{
		"loopback": {127, 0 , 0, 1} ,
		"googleDNS": {8, 8, 8, 8},
	}

	for name , ip := range hosts {
		fmt.Printf("%v :%v \n", name , ip)
	}
	
}

//IPAddr is a data type dor the date
//Printf   is for output the format 

// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
/*
for the example  this {1 , 2 , 3}
to this:    1 2 3 4
*/