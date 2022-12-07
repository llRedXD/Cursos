// - São structs sem identificadores.
// - x := struct { name type }{ name: value }
package main

import "fmt"

func main() {

	x := struct {
		nome string
		idade int
	}{
		nome: "Gabriel",
		idade: 18,
	}
	fmt.Println(x.nome)

}