package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	resultado := Adiciona(2, 2)
	esperado := 4

	if esperado != resultado {
		t.Errorf("Resultado '%d', Esperado '%d'", resultado, esperado)
	}

}
func ExampleAdiciona() {
	soma := Adiciona(1, 4)
	fmt.Println(soma)
	// Output: 5
}
