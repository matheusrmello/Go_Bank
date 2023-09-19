package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	}else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDaMaria := ContaCorrente{}
	contaDaMaria.titular = "Maria"
	contaDaMaria.saldo = 500


	fmt.Println(contaDaMaria.saldo)


	fmt.Println(contaDaMaria.sacar(400))

	fmt.Println(contaDaMaria.saldo)


}
