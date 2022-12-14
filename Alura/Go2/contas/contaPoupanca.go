package contas

import (
	"strconv"
	"github.com/llRedXD/Go2/titular"
)



type ContaPoupanca struct {
	Titular cliente.Titular
	Agencia, Conta, Operacao int
	saldo   float64
}

func (c *ContaPoupanca) Sacar(valor int) string {
	podeSacar := valor >= 0 && valor <= int(c.saldo)
	if podeSacar {
		c.saldo = c.saldo - float64(valor)
		return "Saque realizado com sucesso"
	} else {
		return "Não é possível Sacar esse valor:" + strconv.FormatInt(int64(valor), 10)
	}
}

func (c *ContaPoupanca) Depositar(valor int) (string, int) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += float64(valor)
		return "Deposito Realizado com Sucesso", int(c.saldo)
	} else {
		return "Valor não pode Depositado", int(c.saldo)
	}
}

func (c *ContaPoupanca) Transferir(c2 *ContaPoupanca, valor int) (string, int) {
	podeTransferir := valor >= 0 && valor <= int(c.saldo)
	if podeTransferir {
		c.Sacar(valor)
		c2.Depositar(valor)
		return "Valor transferido com sucesso", int(c.saldo)
	} else {
		return "Não foi possível fazer a transferência com esse valor", valor
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) PagarBoleto(valor int) (string, int) {
	podeTransferir := valor >= 0 && valor <= int(c.saldo)
	if podeTransferir {
		c.Sacar(valor)
		return 	"Boleto Pago com sucesso", int(c.saldo)
	} else {
		return "Erro ao tentar pagar boleto", int(c.saldo)
	}
}
