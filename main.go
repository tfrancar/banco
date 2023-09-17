package main

import (
	"fmt"

	"github.com/tfrancar/banco/clientes"
	"github.com/tfrancar/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {

	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteThiago := clientes.Titular{
		Nome:      "Thiago",
		CPF:       "123.456.789-00",
		Profissao: "Desenvolvedor",
	}

	contaDoThiago := contas.ContaCorrente{
		Titular:       clienteThiago,
		NumeroAgencia: 1234,
		NumeroConta:   253687,
	}

	contaPoupacaDoThiago := contas.ContaPoupanca{
		Titular:       clienteThiago,
		NumeroAgencia: 1234,
		Operacao:      13,
		NumeroConta:   2536879,
	}

	contaDoThiago.Depositar(15000)
	contaPoupacaDoThiago.Depositar(5000)
	PagarBoleto(&contaDoThiago, 14000)
	fmt.Println(contaDoThiago.ObterSaldo())
	//fmt.Println(contaPoupacaDoThiago.ObterSaldo())

}
