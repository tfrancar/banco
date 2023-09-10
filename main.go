package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoThiago := ContaCorrente{
		titular:       "Thiago",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.62,
	}

	// Quando vamos popular todos os campos da estrura podemos usar este segundo modo.
	contaDaKelma := ContaCorrente{
		"Kelma",
		589,
		1234567,
		1125.62,
	}

	// Quando vamos popular campos especificos.
	contaDoGuilherme := ContaCorrente{
		titular: "Guilherme",
		saldo:   1125.62,
	}

	fmt.Println(contaDoThiago)
	fmt.Println(contaDaKelma)
	fmt.Println(contaDoGuilherme)
}
