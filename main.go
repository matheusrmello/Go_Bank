package main

import (
	"fmt"
	"go_bank/contas"
)

func main() {
	contaCleber := contas.ContaPoupanca{}
	contaCleber.Depositar(100)
	PagarBoleto(&contaCleber, 70)


	fmt.Println(contaCleber.Obtersaldo())

	contaLuiza := contas.ContaCorrente{}
	contaLuiza.Depositar(400)
	PagarBoleto(&contaLuiza, 150)

	fmt.Println(contaLuiza.ObterSaldo())

}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
