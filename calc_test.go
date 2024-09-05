package main

import "testing"

func ShouldSomaCorrect(t *testing.T) { //convensão de nomes ShouldSumCorrect (assinatura do método)
	//arrange
	teste := soma(10, 50, 15)
	//act
	resultado := 6
	//assert
	if teste != resultado { //assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func ShouldSomaIncorrect(t *testing.T) { //ShouldSumIncorrect
	teste := soma(10, 50, 15) //arrange
	resultado := 7            //act
	if teste != resultado {   //assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func ShouldMultiCorret(t *testing.T) {
	teste := multiplicar(5, 20)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func ShouldMultiIncorret(t *testing.T) {
	teste := multiplicar(5, 20)
	resultado := 2560
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func ShouldSubtraCorret(t *testing.T) {
	teste := subtrair(88, 100)
	resultado := -5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func ShouldSubtraIncorret(t *testing.T) {
	teste := subtrair(88, 100)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func ShouldDivisaoCorret(t *testing.T) {
	teste := divisao(10)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func ShouldDivisaoIncorret(t *testing.T) {
	teste := divisao(10)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
