// - Callback: passe uma função como argumento a outra função.
package main

import "fmt"


func main() {
    // []int{10,456,78,15,5,68,9,41,12,4,56,78,45,23,}...
    fmt.Println(somaMulti3(soma,[]int{1,2,3,4,5,6,7,8,9}...))
}

func somaMulti3(s func (x ... int) int,n ... int) int {
    multi := []int{}
    for _,v := range n{
        if v % 3 == 0{
            multi = append(multi,v)
        }
    }
    total := soma(multi...)
    return total 
}

func soma(x ... int) int {
    total := 0
    for _,v := range x {
        total += v
    }
    return total
}