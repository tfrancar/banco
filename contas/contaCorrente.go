package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	validarDeposito := valorDoDeposito > 0
	if validarDeposito {
		c.Saldo += valorDoDeposito
	} else {
		return "Depósito não pode ser realizado.", c.Saldo
	}
	return "Depósito realizado com sucesso.  Seu novo saldo é: ", c.Saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64) {

	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso. Seu saldo atual é de ", c.Saldo
	} else {
		return "saldo insuficiente. Seu saldo atual é de:", c.Saldo
	}

}
