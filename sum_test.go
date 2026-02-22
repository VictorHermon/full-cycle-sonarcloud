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

func TestMain(t *testing.T) {

	main()
}
