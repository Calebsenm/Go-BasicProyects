
package test 

import (
  "testing"
  "app/app"
)

func TestSuma ( t *testing.T){
  
  var valor1_ int  = 2
  var valor2_ int  = 2

  suma := app.Suma(valor1_ , valor2_);
  resta := app.Resta(valor1_ , valor2_);

  if suma != 4{
    t.Error("Error ");
  }

  if resta != 0{
    t.Error("Error ");
  }

}
