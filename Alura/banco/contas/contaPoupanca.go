package contas

import (
	cliente "github.com/fabiossilva00/Go/Alura/banco/clientes"
)

type ContaPoupanca struct {
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente %s"
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Valor depositado", c.saldo
	}
	return "Valor incorreto para depositos", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
