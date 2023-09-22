package main

import (
	"fmt"
	"go_bank/contas"
)

func main() {
	contaDaMaria := contas.ContaCorrente{Titular: "Maria", Saldo: 400}
	contaDoMatheus := contas.ContaCorrente{Titular: "Matheus", Saldo: 500}

	status := contaDoMatheus.Transferir(-600, &contaDaMaria)
	fmt.Println(status)

	fmt.Println(contaDaMaria)
	fmt.Println(contaDoMatheus)
}
