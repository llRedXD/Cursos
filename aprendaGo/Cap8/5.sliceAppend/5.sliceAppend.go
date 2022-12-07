// - Effective Go: append (package builtin)
// - x = append(slice, ...values)
// - x = append(slice, slice...)
// - Todd: unfurl → desdobrar, desenrolar
// - Nome oficial: enumeration
package main

import "fmt"

func main() {
	slice := []int {1,2,3,4}
	slice2 := []int {8,9,10}

	slice = append(slice, 5,6,7)
	fmt.Println(slice)
	slice = append(slice, slice2...)
	fmt.Println(slice)

}