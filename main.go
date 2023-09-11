package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDoThiago := ContaCorrente{}
	contaDoThiago.titular = "Thiago"
	contaDoThiago.numeroAgencia = 1106
	contaDoThiago.numeroConta = 147258
	contaDoThiago.saldo = 500.00

	fmt.Println(contaDoThiago.Sacar(600))

	fmt.Println("Saldo atual", contaDoThiago.saldo)
}
