// - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
package main

import "fmt"

func main() {

    defer fmt.Println("Essa era pra ser a primeira")
    fmt.Println("Essa era pra ser a segunda")
    fmt.Println("Essa era pra ser a terceira")
    fmt.Println("Essa era pra ser a quarta")
    
}