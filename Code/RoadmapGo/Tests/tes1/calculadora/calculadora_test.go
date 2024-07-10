

package calculadora


import "testing"


func TestSuma(t *testing.T) {

	valor := Suma(7, 23)

	if valor != 30 {
	  t.Error("Se esperaba 30 y se obtuvo", valor);
	}
}

func TestResta(t *testing.T){

  valor := Resta(2 , 4);
  if valor != -2{
    t.Error("Se esperaba 2 y se obtuvo", valor);
  }
}
