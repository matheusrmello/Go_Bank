package main

import (
	"fmt"

	// "go_bank/clientes"
	"go_bank/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.ObterSaldo()
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}
