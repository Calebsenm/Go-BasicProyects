package main

import (
	"fmt"
	"math/rand"
	"os"

	"os/exec"
)

var (
	Board = [20][20]string{}

	Y = 0
	X = 0

	YY = rand.Intn(19)
	XX =  rand.Intn(19)

	Latest_number = 1
	Decision = 1

	Historial = [][] string {}
	Body = [][] string {}
)

func main() {
	// fill the matriz

	for i := 0; i < len(Board); i++ {

		for j := 0; j < len(Board); j++ {
			Board[i][j] = " "
		}
	}

	iterator := 0

	for {
		iterator++

		cmd := exec.Command("cmd","/c","cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		
		Board[Y][X] = "0"
		Board[YY][XX] = "+"

		for i := 0; i < len(Board); i++ {

			for j := 0; j < len(Board); j++ {
				fmt.Print(Board[i][j])
			}
			fmt.Println()
		}

		fmt.Println("Input the position -> 1 UP 2Down 3Rignt 4Left -> ")
		fmt.Scan(&Decision)


	




	}

}
