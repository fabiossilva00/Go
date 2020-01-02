package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Eu")
	esperado := "Olá, Eu"

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
