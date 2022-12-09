// - É uma maneira de encriptar senhas utilizando hashes.
// - x/crypto/bcrypt
//     - GenerateFromPassword
//     - CompareHashAndPassword
package main

import "fmt"
import "golang.org/x/crypto/bcrypt"

func main() {

    senha := "Teste123456"
    senha1 := "sadasdasdas"

    sb, err := bcrypt.GenerateFromPassword([]byte(senha),5)
    if err != err {
        fmt.Println(err)
    }
    fmt.Println(string(sb))

    t := bcrypt.CompareHashAndPassword(sb, []byte(senha1))
    if t != nil {
        fmt.Println("Erro: ", t)
    } else {
        fmt.Println("Senha Correta")
    }
}