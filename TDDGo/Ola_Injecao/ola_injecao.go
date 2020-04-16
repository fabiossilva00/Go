package main

import (
	"fmt"
	"io"
	"net/http"
)

//Cumprimenta -
func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

//HandlerMeuCumprimento -
func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "Mundo ?")
}

func main() {
	// Cumprimenta(os.Stdout, "Vixi")
	err := http.ListenAndServe(":3000", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		fmt.Println(err)
	}

}
