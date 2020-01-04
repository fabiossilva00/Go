package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagem := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Diz Olá para as pessoas em portugues", func(t *testing.T) {
		resultado := Ola("Eu", "pt")
		esperado := "Olá, Eu"
		verificaMensagem(t, resultado, esperado)
	})

	t.Run("Diz olá para as pessoas em espanhol", func(t *testing.T) {
		resultado := Ola("Eu", "es")
		esperado := "Holá, Eu"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("Diz olá para as pessoas em francês", func(t *testing.T) {
		resultado := Ola("Eu", "fr")
		esperado := "Bonjour, Eu"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("Diz olá para as pessoas em inglês", func(t *testing.T) {
		resultado := Ola("Eu", "en")
		esperado := "Hello, Eu"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("Diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagem(t, resultado, esperado)
	})
}
