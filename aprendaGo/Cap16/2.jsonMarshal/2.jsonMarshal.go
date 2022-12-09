package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
    Nome string 
    Idade int
    Profissao string
}

func main() {

    james := pessoa{"James", 40, "Agente Secreto"}
     
    j, err := json.Marshal(james)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(string(j))

    darth := pessoa{"Darth", 33, "Vilão "}
    d, err := json.Marshal(darth)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(d))

}