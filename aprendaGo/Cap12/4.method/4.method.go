// - Um método é uma função anexada a um tipo.
// - Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
// - Pode-se anexar uma função a um tipo utilizando seu receiver.
// - Utilização: valor.método()
// - Exemplo: o tipo "pessoa" pode ter um método oibomdia()
package main

type pessoa struct {
	nome string
	idade int
}

func (p pessoa) oibomdia() {
	println(p.nome," Oi, bom dia!")
}

func main() {
	claudia := pessoa{"Claudia", 30}
	claudia.oibomdia()
}