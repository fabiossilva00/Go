package olainjecao

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Eitia")

	resultado := buffer.String()
	esperado := "Olá, Eitia"

	if resultado != esperado {
		t.Errorf("Resultado: '%s', Esperado: '%s'", resultado, esperado)
	}
}
