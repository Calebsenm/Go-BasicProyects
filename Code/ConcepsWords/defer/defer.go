
package main


import(
  "fmt"
)

func main (){

  defer fmt.Println("Hola  ")
  fmt.Println(" Hola :) ")
}


/*
Defer
A defer statement defers the execution 
of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, 
but the function call is not executed until the surrounding function returns.


Uno de los principales usos de una instrucción defer es el de limpiar recursos como archivos abiertos, conexiones de red y controladores de bases de datos. Cuando su programa termine con estos recursos, es importante cerrarlos para evitar agotar los límites del programa y permitir que otros programas accedan a esos recursos. defer aporta más claridad a nuestro código y reduce su propensión a experimentar errores mediante la conservación de las invocaciones para cerrar archivos y recursos cerca de las invocaciones abiertas.

*/
