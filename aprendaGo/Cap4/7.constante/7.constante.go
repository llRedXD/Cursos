// - São valores imutáveis.
// - Podem ser tipadas ou não:
//     - const oi = "Bom dia"
//     - const oi string = "Bom dia"
// - As não tipadas só terão um tipo atribuido a elas quando forem usadas.
//     - Ex. qual o tipo de 42? int? uint? float64?
//     - Ou seja, é uma flexibilidade conveniente.
// - Na prática: int, float, string.
//     - const x = y
//     - const ( x = y )

package main

import "fmt"

const x = 10
var y = 10

var c int 
var d float64

func main() {
	// O tipo só é atribuído quando a constante é usada.
	d = x // A constante pode ser atribuída a uma variável de qualquer tipo pois ela não tem tipo definido.
	fmt.Printf("%v %T\n", c, c)
	// d = y Já a variável tem um tipo definido então ela só pode ser atribuída a uma variável do mesmo tipo.

}