package contas

import "github.com/tfrancar/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	validarDeposito := valorDoDeposito > 0
	if validarDeposito {
		c.saldo += valorDoDeposito
	} else {
		return "Depósito não pode ser realizado.", c.saldo
	}
	return "Depósito realizado com sucesso.  Seu novo saldo é: ", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() (string, float64) {
	return "Seu saldo atual é de R$", c.saldo

}
