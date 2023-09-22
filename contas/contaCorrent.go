package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "O valor do deposito eh menor do que zerp", c.Saldo
	}

}
func main() {
	contaDaMaria := ContaCorrente{Titular: "Maria", Saldo: 400}
	contaDoMatheus := ContaCorrente{Titular: "Matheus", Saldo: 500}

	status := contaDoMatheus.Transferir(-600, &contaDaMaria)
	fmt.Println(status)

	fmt.Println(contaDaMaria)
	fmt.Println(contaDoMatheus)
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
