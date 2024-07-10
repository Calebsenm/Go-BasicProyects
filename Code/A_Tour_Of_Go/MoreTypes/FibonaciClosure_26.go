package main

import (
	"fmt"
)

func fibonaci() func() int {
	a := 0
	b := 1
	c := 0

	return func() int {
		a = b
		b = c
		c = a + b

		return b
	}
}

func main() {
	f := fibonaci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
