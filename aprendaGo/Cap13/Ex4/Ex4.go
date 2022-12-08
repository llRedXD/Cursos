// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.
package main

import "fmt"

type pessoa struct {
    nome string
    sobrenome string
    idade int
}

func (p pessoa) identidade() {
    fmt.Println("Minha identidade é:")
    fmt.Println("Nome Completo:", p.nome, p.sobrenome)
    fmt.Println("Idade:", p.idade)
}

func main() {
    gabriel := pessoa{"Gabriel","Neira", 10}
    gabriel.identidade()
}