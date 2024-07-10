
package class 


type Dog struct{

  Color string
  Breed string
}

func ( e * Dog ) NewDog (color , breed string ){
  e.Color = color 
  e.Breed = breed 
}

func (e * Dog ) GetColor() string {
  return e.Color 
} 

func ( e * Dog ) GetBreed () string {
  return e.Breed 
}


