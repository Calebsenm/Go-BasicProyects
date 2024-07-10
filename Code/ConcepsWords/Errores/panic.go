// panic(err)

package main


func employee(name *string, age int){
  if age > 65{
    panic("Age cannot be greater than retirement age")
  }

}

func main() {
  empName := "Samia"
  age := 66

  employee(&empName, age)  

    
}