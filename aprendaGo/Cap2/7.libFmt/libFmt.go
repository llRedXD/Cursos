package main

import "fmt"

// - Setup: strings, ints, bools.
// - Strings: interpreted string literals vs. raw string literals.
//     - Rune literals: cada caractere é um rune literal.
//     - Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte. 
// - Format printing: documentação.
//     - Grupo #1: Print → standard out
//         - func Print(a ...interface{}) (n int, err error)
//         - func Println(a ...interface{}) (n int, err error)
//         - func Printf(format string, a ...interface{}) (n int, err error)
//             - Format verbs. (%v %T)
//     - Grupo #2: Print → string, pode ser usado como variável
//         - func Sprint(a ...interface{}) string
//         - func Sprintf(format string, a ...interface{}) string
//         - func Sprintln(a ...interface{}) string
//     - Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
//         - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
//         - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
//         - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

func main() {
	fmt.Print("Interpreted Vs Literal\n")
	x := "Olá Bom Dia\n \tTudo Bem?" // interpreted string literal
	y := `"Olá Bom Dia\n \tTudo Bem?"` // raw string literal
	fmt.Println(x)
	fmt.Println(y)
	fmt.Print("\n")

	Grupo1()
	Grupo2()
}

func Grupo1() {
	fmt.Print("Grupo 1\n")
	fmt.Print("Print\n")
	fmt.Println("Println")
	fmt.Printf("Printf: %v %T\n\n", "string", "string")
}

func Grupo2() {
	fmt.Print("Grupo 2\n")
	x := "Olá"
	y := "Bom Dia"
	z := fmt.Sprint(x, y)
	t := fmt.Sprintln(x, y)
	w := fmt.Sprintf("%v %v", x, y)

	fmt.Println(z)
	fmt.Println(t)
	fmt.Println(w)
}