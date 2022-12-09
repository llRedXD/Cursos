package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type decodeJson struct{
    Nome string
    Idade int
    Profissao string 
}

func main() {

    jsonCodede := []byte(`{"Nome":"Darth","Idade":33,"Profissao":"Vilão "}`) // aqui fazemos a conversão de string para byte usamos o ` para não ter que fazer um escape de aspas
    fmt.Println(string(jsonCodede))
    darth := decodeJson{} // aqui criamos uma variável do tipo decodeJson
    err := json.Unmarshal(jsonCodede, &darth) // aqui fazemos a conversão de json para struct, para isso passamos o jsonCodede e a variável que vai receber o json sendo um ponteiro
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(darth)

    encoder := json.NewEncoder(os.Stdout) // aqui criamos um encoder que vai escrever no terminal
    encoder.Encode(darth) // aqui fazemos a conversão de struct para json

}