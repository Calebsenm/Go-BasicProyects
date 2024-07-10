package main

import (
	"fmt"
	"os"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Uui Quieto no puedes calcular la raiz cuadrada de un numero  Negativo: %v ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x , ErrNegativeSqrt(x)
	}
	
	z := float64(0.1)

	for {
		z -= (z*z - x) / (2 * z)
		sqrtData := z * z
		fmt.Println(z, " = ", sqrtData)
		if sqrtData == x  || sqrtData <= x {
			return z , nil 
		} 
	}

}

func main() {
	var Data int 
	fmt.Print("Ingresa un numero para calcular la raiz Cuadrada - > ")
	fmt.Scan(&Data)
	data  , err := Sqrt(float64(Data))
	
	if err != nil{
		fmt.Println(err)
		os.Exit(0)		
	}	

	fmt.Println("La raiz Cuadrada de ", Data , " Es: " , data )
}

