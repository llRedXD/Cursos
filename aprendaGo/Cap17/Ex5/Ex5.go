// - Partindo do código abaixo, ordene os []user por idade e sobrenome.
// - Os valores no campo Sayings devem ser ordenados também, e demonstrados de maneira esteticamente harmoniosa.
package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type porIdade []user 

func (a porIdade) Len() int           { return len(a) }
func (a porIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a porIdade) Less(i, j int) bool { return a[i].Age < a[j].Age }

type porSobrenome []user 

func (a porSobrenome) Len() int           { return len(a) }
func (a porSobrenome) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a porSobrenome) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
    
    fmt.Println("Ordenado por Idade")
    sort.Sort(porIdade(users))
    fmt.Println(users)

    mostraSayings(users)

    fmt.Println("Ordenado por Sobrenome")
    sort.Sort(porSobrenome(users))
    fmt.Println(users)

    mostraSayings(users)

}

func mostraSayings(users []user) {
	for i, v := range users {
		fmt.Println(i+1, "\t Nome:", v.First, "\t Sobrenome:", v.Last, "\t Idade:", v.Age)
		for _, m := range v.Sayings {
			fmt.Println("\t\t", m)
		}
	}
}
