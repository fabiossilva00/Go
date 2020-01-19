package carteira

import "testing"

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		esperado := Bitcoin(10)
		confirmaCarteira(t, carteira, esperado)
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{Bitcoin(20)}
		err := carteira.Retirar(10)
		esperado := Bitcoin(10)
		confirmaCarteira(t, carteira, esperado)
		confirmaErroInexistente(t, err)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		err := carteira.Retirar(Bitcoin(1000))
		confirmaCarteira(t, carteira, saldoInicial)
		confirmaErro(t, err, ErroSaldoInsuficiente)
	})
}

func confirmaCarteira(t *testing.T, carteira Carteira, valorEsperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()
	if resultado != valorEsperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, valorEsperado)
	}
}

func confirmaErro(t *testing.T, err error, valorEsperado error) {
	t.Helper()
	if err == nil {
		t.Fatal("Não ocorreu erro")
	}

	if err != valorEsperado {
		t.Errorf("Resultado '%s', Esperado '%s'", err, valorEsperado)
	}
}

func confirmaErroInexistente(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Ocorreu o error '%s'", err)
	}
}
