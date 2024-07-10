package main

import (
    "fmt"
)
func split( sum int ) ( x, y int){
    x = sum + 2
    y = sum - 2
    return
}

func main(){
    fmt.Println(split(10));
    a , b := split(2);

    fmt.Println(a , b);
}


/*
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.
*/
