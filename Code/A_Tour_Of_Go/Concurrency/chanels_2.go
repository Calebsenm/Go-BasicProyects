package main 

import "fmt"

func sum(s []int , c chan int ){
	sum := 0

	for _ , v := range s {
		sum += v
	}

	c <- sum // Send sum to c 
}

func main(){
	s := []int {2 , 6 , 5 , 7, 4 ,9 , 3}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x + y)
}
