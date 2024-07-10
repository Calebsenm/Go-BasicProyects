package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person ) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Caleb", 18}
	b := Person{"Belac", 81}

	fmt.Println(a, b)
}
