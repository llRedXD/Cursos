// - Crie um struct "erroEspecial" que implemente a interface builtin.error.
// - Crie uma função que tenha um valor do tipo error como parâmetro.
// - Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
package main

import "fmt"

type erroEspecial struct {
	msg string
}

func (e erroEspecial) Error() string {
	return "Erro Especial"
}

func erros(e error) {
	fmt.Printf("Erro: após a execução ocorreu o erro:%v e %v", e, e.(erroEspecial).msg )
}

func main() {

	erro := erroEspecial{
		msg: "Erro Tosco",
	}
	erros(erro)

}
