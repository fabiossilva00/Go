package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{
		"teste": "Isso é um teste",
	}

	t.Run("Palavra Conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "Isso é um teste"

		comparaString(t, resultado, esperado)
	})

	t.Run("Palavra Desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecida")

		comparaError(t, err, ErrNaoEncontrado)
	})

}

func TestAdd(t *testing.T) {
	dicionario := Dicionario{}
	palavra := "teste"
	definicao := "Isso é apenas um teste"
	err := dicionario.Add(palavra, definicao)
	comparaError(t, err, nil)
	comparacaoDefinicao(t, dicionario, palavra, definicao)

	t.Run("Palavra Existente", func(t *testing.T) {
		palavraExistente := "teste"
		definicaoExistente := "Isso é apenas um teste"
		dicionario := Dicionario{palavraExistente: definicaoExistente}
		err := dicionario.Add(palavraExistente, definicaoExistente)
		comparaError(t, err, ErrPalavraExistente)
	})

}

func TestAtualiza(t *testing.T) {

	t.Run("Palavra Existente", func(t *testing.T) {
		palavraExistente := "teste"
		definicaoExistente := "Isso é apenas mais um teste"
		dicionarioExistente := Dicionario{palavraExistente: definicaoExistente}
		novaDefinicaoExistente := "Isso é apenas um teste modificado"

		err := dicionarioExistente.Atualiza(palavraExistente, novaDefinicaoExistente)

		comparaError(t, err, nil)
		comparacaoDefinicao(t, dicionarioExistente, palavraExistente, novaDefinicaoExistente)
	})

	t.Run("Palavra Nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavraNova := "teste"
		definicaoNova := "Isso é apenas um teste novo"

		err := dicionario.Atualiza(palavraNova, definicaoNova)

		comparaError(t, err, ErrPalavraInexistente)
	})
}

func TestDelete(t *testing.T) {
	palavra := "teste"
	definicao := "Isso é apenas um teste deletável"
	dicionario := Dicionario{palavra: definicao}

	dicionario.Delete(palavra)

	_, err := dicionario.Busca(palavra)

	if err != ErrNaoEncontrado {
		t.Errorf("A palavra: '%s' não foi deletada", palavra)
	}

}

func comparacaoDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)

	if err != nil {
		t.Fatal("Deveria ter encontrado uma palavra", err)
	}

	comparaString(t, resultado, definicao)
}

func comparaError(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}

func comparaString(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
