
/*
In this example, we shortened
x int, y int
to
x, y int
*/

package main 

import (
    "fmt"
)

func add( a , b int ) int {
    return a + b;
}

func main(){
    fmt.Println("The add is ", add(2 , 4) );
}
