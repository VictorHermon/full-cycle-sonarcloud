package main

import "testing"

func TestSoma(t *testing.T) {

	resultado := Soma(3, 4)

	if resultado != 7 {
		t.Errorf("Resultado = %d; esperado 7", resultado)
	}
}

func TestSub(t *testing.T) {

	resultado := sub(10, 5)

	if resultado != 5 {
		t.Errorf("Resultado = %d; esperado 5", resultado)
	}
}

func TestMult(t *testing.T) {

	resultado := mult(6, 7)

	if resultado != 42 {
		t.Errorf("Resultado = %d; esperado 42", resultado)
	}
}

func TestDiv(t *testing.T) {

	resultado := div(20, 4)

	if resultado != 5 {
		t.Errorf("Resultado = %d; esperado 5", resultado)
	}
}
