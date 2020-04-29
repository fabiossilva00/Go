package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestContagem(t *testing.T) {
	t.Run("Reveillon!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `10
9
8
7
6
5
4
3
2
1
Feliz Ano Novo!!!`

		if resultado != esperado {
			t.Errorf("\nResultado %s, \nEsperado %s", resultado, esperado)
		}
	})

	t.Run("Pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}

		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("Esperado %v chamadas, Resultado %v chamadas", esperado, spyImpressoraSleep.Chamadas)
		}
	})
}

func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second
	tempoSpy := &TempoSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
	sleeper.Pausa()

	if tempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("Deveria ter pausado por %v, mas parou por %v", tempoPausa, tempoSpy.duracaoPausa)
	}

}
