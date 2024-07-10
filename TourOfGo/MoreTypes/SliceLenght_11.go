package main 

import (
    "fmt"
)

func printSlice(s []int){
    fmt.Printf("len = %d cap= %d %v \n ",len(s) , cap(s) , s)
}

func main(){
    s := []int {2 , 3 , 5 , 7, 11 , 13}
    printSlice(s)

    //slice the slice to give it zero lenth
    s = s[:0]
    printSlice(s)

    //Extens its length
    s = s[:4]
    printSlice(s)
    
    //Drop its first two values 
    s = s[2:]
    printSlice(s)
}
