// - Em Go, valores podem ter mais que um tipo.
// - Uma interface permite que um valor tenha mais que um tipo.
// - Declaração: keyword identifier type → type x interface
// - Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
// - Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
// - Esse tipo será o seu tipo e também o tipo da interface.

// - Exemplos:
//   - Os tipos profissão1 e profissão2 contem o tipo pessoa
//   - Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
//   - Implementam a interface gente
//   - Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
//   - Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:
//     switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }*
//   - Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
//     https://play.golang.org/p/zGKr7cvTPF
//   - Go Playground ao vivo:
//     https://play.golang.org/p/njiKbTT20Cr
//
// - Onde se utiliza?
//   - Área de formas geométricas (gobyexample.com)
//   - Sort
//   - DB
//   - Writer interface: arquivos locais, http request/response
//
// - Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.
package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

type doutor struct {
	pessoa
	especialidade string
	salario float64
}

type programador struct {
	pessoa
	estresse int
	salario float64
}

func (p doutor) oibomdia() {
	fmt.Println("oi bom dia", p.nome, "sou doutor")
}

func (p programador) oibomdia() {
	fmt.Println("oi bom dia", p.nome, "sou programador estou com ", p.estresse, "%  de estresse" )
}

type gente interface {
	oibomdia()
}

func serHumano(g gente) {
	g.oibomdia()
	switch g.(type) {
		case doutor:
			fmt.Println("Eu como doutor ganho", g.(doutor).salario)
	
		case programador:
			fmt.Println("Eu como programador ganho", g.(programador).salario)
	}
}

func main() {
	doctor := doutor{pessoa{"joao", 30}, "cardiologista", 10000}
	programmer := programador{pessoa{"maria", 25}, 100, 10000}
	// doctor.oibomdia() // Posso chamar o método diretamente
	// programmer.oibomdia() 
	serHumano(doctor) // Posso chamar o método através da interface (poliformismo) não importa o tipo
	serHumano(programmer)
}