package main 

import (
    "fmt"
)

func main(){
    sum := 0

    for i := 0; i < 10; i++{
        sum += i
        fmt.Println("Sum is : ", sum ," in ", i )
    }
    fmt.Println(sum)
}

//Go has only one looping construct, the for loop.
