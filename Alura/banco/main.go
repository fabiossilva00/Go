package main

import (
	//clientes "github.com/fabiossilva00/Go/Alura/banco/clientes"
	"fmt"

	contas "github.com/fabiossilva00/Go/Alura/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaUserPoupa := contas.ContaPoupanca{}
	contaUserPoupa.Depositar(250)
	fmt.Println(contaUserPoupa.ObterSaldo())
	PagarBoleto(&contaUserPoupa, 50)
	fmt.Println(contaUserPoupa)

	contaUserCorre := contas.ContaCorrente{}
	fmt.Println(contaUserCorre)
}

//Estudo variavel private get
// func main() {
// 	contaExemplo := cC.ContaCorrente{}
// 	contaExemplo.Depositar(1000)
// 	fmt.Println(contaExemplo)
// 	fmt.Println(contaExemplo.ObterSaldo())
// }

//Estudo Structs composiçao
// func main() {
// 	clinteUser := clientes.Titular{Nome: "User", CPF: "123.123"}
// 	contaCorrenteUser := cC.ContaCorrente{NumeroConta: 0123, NumeroAgencia: 001, Titular: clinteUser}

// 	fmt.Println(contaCorrenteUser)
// }

//Estudo Pacotes
// func main() {
// 	contaUser1 := cC.ContaCorrente{Titular: "User 1", Saldo: 300}
// 	contaUser2 := cC.ContaCorrente{Titular: "User 2", Saldo: 100}
// 	fmt.Println(contaUser1.Saldo)
// 	fmt.Println(contaUser2.Saldo)

// 	status := contaUser1.Tranferencia(200, &contaUser2)

// 	fmt.Println(status)
// 	fmt.Println(contaUser1.Saldo)
// 	fmt.Println(contaUser2.Saldo)

// }

//Estudo de funções
// func main() {
// 	contaUser := ContaCorrente{}
// 	contaUser.titular = "User"
// 	contaUser.saldo = 1000.00

// 	fmt.Println(contaUser.saldo)

// 	fmt.Println(contaUser.Sacar(-10000))

// 	status, valor := contaUser.Depositar(100)
// 	fmt.Println(contaUser)
// 	fmt.Println(status)
// 	fmt.Println(valor)

// }

//Estudo de variaveis e Comparações
// func main() {

// 	contaUser := ContaCorrente{
// 		titular:       "User",
// 		numeroAgencia: 001,
// 		numeroConta:   123456,
// 		saldo:         1.1,
// 	}
// 	contaUsuario := ContaCorrente{
// 		"Usuario", 001, 987654, 9.9,
// 	}

// 	fmt.Println(contaUser == contaUsuario)
// 	fmt.Println(contaUser.numeroAgencia == contaUsuario.numeroAgencia)

// 	contaOutra := ContaCorrente{}
// 	contaOutra.titular = "Outra"
// 	contaOutra.numeroConta = 001
// 	contaOutra.numeroAgencia = 0001
// 	contaOutra.saldo = 100000.99

// 	var contaDepois *ContaCorrente
// 	contaDepois = new(ContaCorrente)

// 	fmt.Println(contaUser)
// 	fmt.Println(contaUsuario)
// 	fmt.Println(contaOutra)
// 	fmt.Println(*contaDepois)
// }
