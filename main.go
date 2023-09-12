package main

import (
	"fmt"

	"github.com/tfrancar/banco.git/contas"
)

func main() {

	contaDoThiago := contas.ContaCorrente{}
	contaDoThiago.Titular = "Thiago"
	contaDoThiago.NumeroAgencia = 1106
	contaDoThiago.NumeroConta = 147258
	contaDoThiago.Saldo = 500.00

	contaDaKelma := contas.ContaCorrente{}
	contaDaKelma.Titular = "Kelma"
	contaDaKelma.NumeroAgencia = 1107
	contaDaKelma.NumeroConta = 147259
	contaDaKelma.Saldo = 1500.00

	status, valor := contaDaKelma.Transferir(200, &contaDoThiago)
	fmt.Println(status, valor)
	fmt.Println("Saldo da conta de", contaDoThiago.Titular, "é de R$", contaDoThiago.Saldo)

	// fmt.Println(contaDoThiago.Sacar(500))
	// saldo, valor := contaDoThiago.Depositar(1000)
	// fmt.Println(saldo, valor)

	// fmt.Println("Saldo atual", contaDoThiago.saldo)
}
