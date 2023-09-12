package contas

import "github.com/tfrancar/banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	validarDeposito := valorDoDeposito > 0
	if validarDeposito {
		c.saldo += valorDoDeposito
	} else {
		return "Depósito não pode ser realizado.", c.saldo
	}
	return "Depósito realizado com sucesso.  Seu novo saldo é: ", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64) {

	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso. Seu saldo atual é de ", c.saldo
	} else {
		return "saldo insuficiente. Seu saldo atual é de:", c.saldo
	}

}

func (c *ContaCorrente) ObterSaldo() (string, float64) {
	return "Seu saldo atual é de R$", c.saldo

}
