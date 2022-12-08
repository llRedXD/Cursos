// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//   - Área do círculo: 2 * π * raio
//   - Área do quadrado: lado * lado
//
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"
package main

import (
	"fmt"
	"math"
)

type quadrado struct {
    lado float64
}

type circulo struct {
    raio float64
}

func (q quadrado) area() float64 {
    return q.lado * q.lado
}

func (c circulo) area() float64 {
    return math.Pi * (c.raio * c.raio)
}

type figura interface {
    area() float64
}

func info(f figura)  {
    switch f.(type) {
        case quadrado :
            fmt.Printf("A área de uma caixa de 6cm é de %vcm \n", f.area())
            f.(quadrado).area()    
        case circulo :
            fmt.Printf("A área de uma bola de 5cm é de %vcm \n", f.area())
        }
    }

func main() {

    caixa := quadrado{6}
    bola := circulo{5}
    info(caixa)
    info(bola)
    
}