package contas

import (
	cliente "github.com/fabiossilva00/Go/Alura/banco/clientes"
)

type ContaCorrente struct {
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente %s"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Valor depositado", c.saldo
	}
	return "Valor incorreto para depositos", c.saldo
}

func (c *ContaCorrente) Tranferencia(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
