package iteracao

import "testing"

import "fmt"

func TestRepetir(t *testing.T) {
	resultado := Repetir("A", 5)
	esperado := "AAAAA"

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}

}

func ExampleRepetir() {
	result := Repetir("A", 5)
	fmt.Println(result)
	// Output: AAAAA
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("A", 5)
	}
}
