// Realizar um programa que:
// - Tenha uma interface para calcular o perímetro e área de uma figura geométrica
// - Tenha uma struct para cada figura geométrica(quadrado, circulo)
package main

import (
	"fmt"
	"math"
)

type figura interface {
	perimetro() float64
	area() float64
}

type pad struct {
	altura float64
	comprimento float64
}

func (p pad) perimetro() float64 {
	return float64(p.altura) + float64(p.comprimento)
}

func (p pad) area() float64 {
	return float64(p.altura) * float64(p.comprimento)
}

type circle struct {
	radius float64
}

func (c circle) perimetro() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main() {

	p := pad{4,3}
	c := circle{5}

	fmt.Println(p.perimetro())
	fmt.Println(c.area())

}