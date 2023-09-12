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
		Saldo:         500.00,
	}

	contaDaKelma := contas.ContaCorrente{
		Titular:       clienteKelma,
		NumeroAgencia: 4321,
		NumeroConta:   876352,
		Saldo:         11500.00,
	}

	fmt.Println(contaDoThiago, contaDaKelma)

}
