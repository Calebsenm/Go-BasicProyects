package main 
import (
    "fmt"
)

func main(){
    sum := 1 

    for sum < 100000 {
        sum += sum
        fmt.Println(sum)
    }
    fmt.Println(sum)
}

//At that point you can drop the semicolons: C's while is spelled for in Go.

