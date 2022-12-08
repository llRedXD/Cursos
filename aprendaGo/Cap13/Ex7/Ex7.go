// - Atribua uma função a uma variável.
// - Chame essa função.
// - Solução: https://play.golang.org/p/RMHLL3N5Ww
package main

import "fmt"

func main() {
    x := 10
    y := func (x int) {
        fmt.Println("O valor de x é", x)
    }
    y(x)
}