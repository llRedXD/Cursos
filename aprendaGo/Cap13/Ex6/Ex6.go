// - Crie e utilize uma função anônima.
package main

import "fmt"

func main() {
    x := 10
    func (x int) {
        fmt.Println("O valor de x é", x)
    }(x)

}