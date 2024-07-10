
package main 

import (
  "fmt"
  "main/class"
)

func main (){

  el := class.Employed{Id : 101}
  el.SetEmployed("Caleb","232323","loro")

  fmt.Println(el.GetId())
  
  

  dog1 :=  class.Dog{"Browm","Pastor Aleman"}
  fmt.Println(dog1.GetColor())

  dog2 := new(class.Dog)
  dog2.NewDog("white","chiguagua")
  theColor :=  dog2.GetColor()
  theBreed :=  dog2.GetBreed()

  fmt.Println(theColor ," - ", theBreed )




}
