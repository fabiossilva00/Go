package carteira

import "fmt"

import "errors"

var ErroSaldoInsuficiente = errors.New("Saldo insuficiente para retiradada")

//Bitcoin -
type Bitcoin int

//Carteira -
type Carteira struct {
	saldo Bitcoin
}

//Stringer - Dados Bitcoin
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Depositar -
func (c *Carteira) Depositar(deposito Bitcoin) {
	c.saldo += deposito
}

//Retirar -
func (c *Carteira) Retirar(saque Bitcoin) error {
	if saque > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= saque
	return nil
}

//Saldo -
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
