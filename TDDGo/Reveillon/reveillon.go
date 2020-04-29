package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const reveillon = "Feliz Ano Novo!!!"
const inicioContagem = 10
const escrita = "escrita"
const pausa = "pausa"

//Sleeper -
type Sleeper interface {
	Pausa()
}

//SleeperSpy -
type SleeperSpy struct {
	Chamadas int
}

//SpyContagemOperacoes -
type SpyContagemOperacoes struct {
	Chamadas []string
}

//SleeperConfiguravel -
type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

//TempoSpy -
type TempoSpy struct {
	duracaoPausa time.Duration
}

//Pausa -
func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

//Pausa -
func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

//Pausa -
func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

//Write -
func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

//Contagem -
func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprintf(saida, reveillon)
}

func main() {
	sleeper := &SleeperConfiguravel{
		1 * time.Second, time.Sleep,
	}

	Contagem(os.Stdout, sleeper)
}
