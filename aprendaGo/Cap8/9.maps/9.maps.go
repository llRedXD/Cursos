// - Utiliza o formato key:value.
// - E.g. nome e telefone
// - Performance excelente para lookups.
// - map[key]value{ key: value }
// - Acesso: m[key]
// - Key sem value retorna zero. Isso pode trazer problemas.
// - Para verificar: comma ok idiom.
//   - v, ok := m[key]
//   - ok é um boolean, true/false
//
// - Na prática: if v, ok := m[key]; ok { }
// - Para adicionar um item: m[v] = value
// - Maps não tem ordem.
package main

import "fmt"

func main() {
	amigos := map[string]int {
		"joao": 123456,
		"maria": 654321,
		"pedro": 987654,
	}

	fmt.Println(amigos)
	
	fmt.Println(amigos["maria"])
	
	amigos["carlos"] = 4444
	
	fmt.Println(amigos)

	if sera, ok := amigos["amiga"]; !ok { // comma ok  serve para verificar se o valor existe para saber se é um valor 0 ou não
		fmt.Println("Não tenho amiga")
	} else {
		fmt.Println(sera)
	}
	

}