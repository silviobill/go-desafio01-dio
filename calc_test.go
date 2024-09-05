package main

import "testing"

func TesteDeSoma(t *testing.T) { //convensão de nomes ShouldSumCorrect (assinatura do método)
	//arrange
	teste := soma(10, 50, 15)
	//act
	resultado := 6
	//assert
	if teste != resultado { //assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TesteDeSoma02(t *testing.T) { //ShouldSumIncorrect
	teste := soma(10, 50, 15) //arrange
	resultado := 7            //act
	if teste != resultado {   //assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TesteDeMultiplicacao(t *testing.T) {
	teste := multiplica(5, 20)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TesteDeMultiplicacao2(t *testing.T) {
	teste := multiplica(5, 20)
	resultado := 2560
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TesteDeSubtracao(t *testing.T) {
	teste := subtrai(88, 100)
	resultado := -5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TesteDeSubtracao2(t *testing.T) {
	teste := subtrai(88, 100)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TesteDeDivisao(t *testing.T) {
	teste := divide(10)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TesteDeDivisao2(t *testing.T) {
	teste := divide(10)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
