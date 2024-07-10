package main

import (
    "fmt"
    "runtime"
)

func main(){
    fmt.Println("Go runs on ")
    
    switch os := runtime.GOOS; os {
    case "darwing":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Printf("%s.\n", os)
    }
}
