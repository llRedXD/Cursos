package main

import (
	"fmt"
	"github.com/llRedXD/Go2/contas"
	"github.com/llRedXD/Go2/titular"
)

type caixa interface {
	Sacar(x int) string
}

func Paga(c caixa, valor int) {
	c.Sacar(valor)
	switch c.(type) {
	case *contas.ContaPoupanca:
		fmt.Println("Dinheiro Retirado da Poupança")
	case *contas.ContaCorrente:
		fmt.Println("Dinheiro Retirado da Conta Corrente")

	}	
}	 

func main() {
	Gabriel := contas.ContaCorrente{
		Titular: cliente.Titular{Nome: "Gabriel",Cpf: 51331384826,Profissao: "Programador"},
		Agencia: 15,
		Conta:   120,
	}
	fmt.Println(Gabriel)

	Bruna := contas.ContaCorrente{
		Titular: cliente.Titular{Nome: "Bruna", Cpf: 23456678958, Profissao: "Bancaria"},
		Agencia: 20,
		Conta:   170,
	}
	fmt.Println(Bruna)

	Cris := new(contas.ContaCorrente)
	Cris.Titular = cliente.Titular{Nome: "Cris",Cpf: 15795712585,Profissao: "Policial"}
	Cris.Agencia = 12
	fmt.Println(Cris)

	fmt.Println("\nBruna")
	fmt.Println(Bruna.Sacar(300))

	fmt.Println("\nGabriel")
	dep, valor := Gabriel.Depositar(300)
	fmt.Println(dep, "Saldo Atual:", valor)

	fmt.Println("\nTransferência")
	msg, valor := Gabriel.Transferir(Cris, 200)
	fmt.Println(msg, "Saldo Atual:", valor)

	fmt.Println("\nSaldo")
	Paga(&Gabriel, 50)
	fmt.Println(Gabriel.ObterSaldo())

	fmt.Println("\nConta Poupança")

	Pedro := contas.ContaPoupanca {
		Titular: cliente.Titular{Nome: "Pedro",Cpf: 45615932420,Profissao: "Pedreiro"},
		Agencia: 50,Conta: 50,Operacao: 1,
	}
	Pedro.Depositar(50)
	Paga(&Pedro, 10)
	fmt.Println(Pedro.ObterSaldo())

}
