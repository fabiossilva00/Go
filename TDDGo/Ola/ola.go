package main

import "fmt"

func main() {
	fmt.Println(Ola("Eu", "pt"))
}

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Holá, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaIngles = "Hello"

// Ola - Comentario
func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	return verificaIdioma(idioma) + nome
}

func verificaIdioma(idioma string) (saudacao string) {
	switch idioma {
	case "es":
		saudacao = prefixoOlaEspanhol
	case "fr":
		saudacao = prefixoOlaFrances
	case "en":
		saudacao = prefixoOlaIngles
	default:
		saudacao = prefixoOlaPortugues
	}
	return
}
