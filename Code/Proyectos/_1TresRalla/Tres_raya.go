package main

import "fmt"
import "os"
import "os/exec"
// the box

func box(A [9]string) string {

  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
	
  fmt.Println("+---+---+---+")
	fmt.Println("| " + A[0] + " | " + A[1] + " | " + A[2] + " |")
	fmt.Println("+---+---+---+")
	fmt.Println("| " + A[3] + " | " + A[4] + " | " + A[5] + " |")
	fmt.Println("+---+---+---+")
	fmt.Println("| " + A[6] + " | " + A[7] + " | " + A[8] + " |")
	fmt.Println("+---+---+---+")
	return ""

}
// the winner
func Winner(B [9]string,X string)string{

	if B[0] == X  && B[1] == X && B[2] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} else if B[3] == X  && B[4] == X && B[5] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} else if B[6] == X  && B[7] == X && B[8] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} else if B[0] == X  && B[3] == X && B[6] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} else if B[1] == X  && B[4] == X && B[7] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} else if B[2] == X  && B[5] == X && B[8] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} else if B[0] == X  && B[4] == X && B[8] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} else if B[2] == X  && B[4] == X && B[6] == X{
		fmt.Print("Felicidades Has ganado el juego")
		os.Exit(0)
	} 


	return "Felicidades has gaando la partida"


}

func main() {

  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
  
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::::::::::::::::::::Welcome to my game::::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")

	var Player_One string
	var Player_Two string
	A := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	B := 0
	C := 0

	counter := 0

	for {
		counter++

		fmt.Print("Please input the name of player one -> ")
		fmt.Scan(&Player_One)
		fmt.Print("Please input the name of player two -> ")
		fmt.Scan(&Player_Two)

		if Player_One == Player_Two {
			fmt.Println("The name is repeated")
		} else {
			break
		}
	}

	for {
		B++
		box(A)

		if C%2 == 0 {
			Winner(A,"O")


			fmt.Println(Player_One + " Te corresponde la letra " + "X")

			H := 0
			fmt.Print("Where do you want to play? -> ")
			fmt.Scan(&H)

			if H <= 9 && H >= 1 {
				
        		if A[H-1] == " "{
          			A[H-1] = "X"

        	} 	else{
          			fmt.Println("La ubicacion está ocupada")
          				C = C + 1
        		}
				
			} else {
				fmt.Println("Error")
			}

		} else {

			Winner(A,"X")
			fmt.Println(Player_Two + " Te corresponde la letra " + "O")

			H := 0
			fmt.Print("Where do you want to play? -> ")
			fmt.Scan(&H)

			if H <= 9 && H >= 1 {
				
				if A[H-1] == " " {
            A[H-1] = "O"
        } else{
          fmt.Println("La ubicacion está ocupada")
          C = C + 1
        }
			} else {
				fmt.Println("Error")
			}
		}
		C++
	}
}
