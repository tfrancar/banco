package main

import (
	"fmt"

	"github.com/tfrancar/banco/clientes"
	"github.com/tfrancar/banco/contas"
)

func main() {

	clienteThiago := clientes.Titular{
		Nome:      "Thiago",
		CPF:       "123.456.789-00",
		Profissao: "Desenvolvedor",
	}

	clienteKelma := clientes.Titular{
		Nome:      "Kelma",
		CPF:       "789.123.456-89",
		Profissao: "Financeiro",
	}

	contaDoThiago := contas.ContaCorrente{
		Titular:       clienteThiago,
		NumeroAgencia: 1234,
		NumeroConta:   253687,
	}

	contaDaKelma := contas.ContaCorrente{
		Titular:       clienteKelma,
		NumeroAgencia: 4321,
		NumeroConta:   876352,
	}

	contaDaKelma.Depositar(1500)
	contaDoThiago.Depositar(500)

	fmt.Println(contaDaKelma.ObterSaldo())
	fmt.Println(contaDoThiago.ObterSaldo())

	fmt.Println(contaDoThiago, contaDaKelma)

}
