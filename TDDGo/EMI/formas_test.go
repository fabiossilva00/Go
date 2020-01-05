package emi

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("Resultado '%f', Esperado '%f'", resultado, esperado)
	}
}

func TestArea(t *testing.T) {

	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("Resultado '%#f', Esperado '%f', Forma '%#v'", resultado, esperado, forma)
		}
	}

	t.Run("De varias formas", func(t *testing.T) {
		testesArea := []struct {
			forma    Forma
			esperado float64
		}{
			{forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
			{forma: Circulo{Raio: 10}, esperado: 314.159},
			{forma: Triangulo{Base: 12.0, Altura: 6.0}, esperado: 36.0},
		}

		for _, tt := range testesArea {
			verificaArea(t, tt.forma, tt.esperado)
		}
	})

	t.Run("De um Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		esperado := 72.0
		verificaArea(t, retangulo, esperado)
	})

	t.Run("De um Circulo", func(t *testing.T) {
		circulo := Circulo{10.0}
		esperado := 314.159
		verificaArea(t, circulo, esperado)
	})

}
