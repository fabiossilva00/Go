package main

import "fmt"

func main() {
	fmt.Println(Ola("Eu"))
}

// Ola - Comentario
func Ola(nome string) string {
	return "Olá, " + nome
}
