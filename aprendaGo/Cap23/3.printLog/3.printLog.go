// - Opções:
//   - fmt.Println() → stdout
//   - log.Println() → timestamp + pode-se determinar onde o erro ficará logado
//   - log.Fatalln() → os.Exit(1) sem defer
//   - log.Panicln() → println + panic → funcões em defer rodam; dá pra usar recover
//   - panic()
//
// - Recomendação: use log.
// - Código:
//   - 1. fmt.Println
//   - 2. log.Println
//   - 3. log.SetOutput
//   - 4. log.Fatalln
//   - 5. log.Panicln
//   - 6. panic
package main

import (
	// "fmt"
	"log"
	"os"
)

func main() {

    f, err := os.Create("log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    // log.SetOutput(f) O setOutput é para onde o log vai ser escrito, no caso o arquivo log.txt

    f2, err := os.Open("no_text-file.txt")
    if err != nil { 
        // fmt.Println("Erro: ", err) // fmt.Println não é recomendado para erros
        log.Println("Erro:", err) // log.Println é recomendado para erros ele imprime o erro e a data e hora
        log.Fatal(err) // log.Fatal é recomendado para erros ele imprime o erro e a data e hora e fecha o programa
        // log.Fatalln(err) // log.Fatalln é recomendado para erros ele imprime o erro e a data e hora e fecha o programa 
        log.Panic(err) // log.Panic é recomendado para erros ele imprime o erro e a data e hora e fecha o programa  mas ele executa o defer
    }
    defer f2.Close()

    


}