package contas

import (
	"strconv"
	"github.com/llRedXD/Go2/titular"
)



type ContaCorrente struct {
	Titular cliente.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valor int) string {
	podeSacar := valor >= 0 && valor <= int(c.saldo)
	if podeSacar {
		c.saldo = c.saldo - float64(valor)
		return "Saque realizado com sucesso"
	} else {
		return "Não é possível Sacar esse valor:" + strconv.FormatInt(int64(valor), 10)
	}
}

func (c *ContaCorrente) Depositar(valor int) (string, int) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += float64(valor)
		return "Deposito Realizado com Sucesso", int(c.saldo)
	} else {
		return "Valor não pode Depositado", int(c.saldo)
	}
}

func (c *ContaCorrente) Transferir(c2 *ContaCorrente, valor int) (string, int) {
	podeTransferir := valor >= 0 && valor <= int(c.saldo)
	if podeTransferir {
		c.Sacar(valor)
		c2.Depositar(valor)
		return "Valor transferido com sucesso", int(c.saldo)
	} else {
		return "Não foi possível fazer a transferência com esse valor", valor
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) PagarBoleto(valor int) (string, int) {
	podeTransferir := valor >= 0 && valor <= int(c.saldo)
	if podeTransferir {
		c.Sacar(valor)
		return 	"Boleto Pago com sucesso", int(c.saldo)
	} else {
		return "Erro ao tentar pagar boleto", int(c.saldo)
	}
}
