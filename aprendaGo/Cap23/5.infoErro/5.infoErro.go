// - Para que nossas funções retornem erros customizados, podemos utilizar:
//     - return errors.New()
//     - return fmt.Errorf() ← tem um errors.New() embutido, olha na fonte!
//     - https://golang.org/pkg/builtin/#error
// - “Error values in Go aren’t special, they are just values like any other, and so you have the entire language at your disposal.” - Rob Pike
// - Código: 
//     - 1. errors.New
//     - 2. var errors.New
//     - 3. fmt.Errorf
//     - 4. var fmt.Errorf
//     - 5. type + method = error interface
package main

import "fmt"

func main() {

    fmt.Println("Hello World")

}