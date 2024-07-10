package main 

import (
	"fmt"
	"os"
	"os/exec"
)

var(
	Player_One string
	Player_Two string
	A = [9] string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	C = 0
	B = 0
)

func box(A[9] string) string{

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
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} else if B[3] == X  && B[4] == X && B[5] == X{
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} else if B[6] == X  && B[7] == X && B[8] == X{
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} else if B[0] == X  && B[3] == X && B[6] == X{
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} else if B[1] == X  && B[4] == X && B[7] == X{
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} else if B[2] == X  && B[5] == X && B[8] == X{
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} else if B[0] == X  && B[4] == X && B[8] == X{
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} else if B[2] == X  && B[4] == X && B[6] == X{
		fmt.Print("Felicidades ", X ," Has ganado el juego")
		os.Exit(0)
	} 

	return "Felicidades has gaando la partida"
}

func Logic(J,I  string) string{
	Winner(A,I)

	fmt.Println(Player_One + " Te corresponde la letra " + "X")

	H := 0
	fmt.Print("Where do you want to play? -> ")
	fmt.Scan(&H)

	if H <= 9 && H >= 1 {
				
        if A[H-1] == " "{
  			A[H-1] = J
					
        } 	else{
  				fmt.Println("La ubicacion está ocupada")
          		C = C + 1
        	}
				
	} 	else {
			fmt.Println("Error")
			C = C + 1
		}
	return ""
}

func main(){
	
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::::::::::::::::::::Welcome to my game::::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")

	counter := 0
	for{
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

	for{
		B++
		box(A)
		if C % 2 == 0{
			Logic("X","O")
		} else{
			Logic("O","X")
		}
		C++
	}

}