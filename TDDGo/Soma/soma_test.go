package soma

import "testing"

import "reflect"

func TestSoma(t *testing.T) {
	t.Run("Teste com Slice", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("Resultado '%d ', Esperado '%d', Dados '%v'", resultado, esperado, numeros)
		}
	})
}

func TestTudo(t *testing.T) {
	resultado := Tudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado '%d', Esperado '%d'", resultado, esperado)
	}
}

func TestTodoResto(t *testing.T) {

	verificaMensagem := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Resultado '%v', Esperado '%v'", resultado, esperado)
		}
	}

	t.Run("Soma de alguns slices", func(t *testing.T) {
		resultado := TodoResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("Soma slice vazios de forma segura", func(t *testing.T) {
		resultado := TodoResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		verificaMensagem(t, resultado, esperado)
	})
}
