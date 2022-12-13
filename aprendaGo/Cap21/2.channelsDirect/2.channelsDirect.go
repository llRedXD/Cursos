// - Canais podem ser direcionais.
// - E isso serve pra...?
// - Um send channel e um receive channel são tipos diferentes. Isso permite que os type-checking mechanisms do compilador façam com que não seja possível, por exemplo, escrever num canal de leitura.
// - Aos aventureitos: https://stackoverflow.com/questions/1...
// - Canais bidirecionals (send & receive)
//     - send chan←
//         - error: "invalid operation: ←cs (receive from send-only type chan← int)"
//     - receive ←chan
//         - error: "invalid operation: cr ← 42 (send to receive-only type ←chan int)"
// - Exemplo: https://play.golang.org/p/TlcSm8bHkW
// - A seta sempre aponta para a esquerda.
// - Assignment/conversion: 
//     - de geral para específico
//     - de específico para geral não
//     - Exemplos:
//         - geral pra específico: https://play.golang.org/p/H1uk4YGMBB
//         - específico pra específico: https://play.golang.org/p/8JkOnEi7-a 
//         - específico pra geral: https://play.golang.org/p/4sOKuQRHq7
//         - atribuição tipos !=: https://play.golang.org/p/bG7H6l03VQ 

// Utilizando canais

// - Em funcs podemos especificar:
//     - receive channel
//         - Parâmetro receive channel: (c ←chan int)
//         - No scope dessa função, esse canal só recebe
//         - Não podemos fechar um receive channel
//     - send channel 
//         - Parâmetro send channel: (c chan← int)
//         - No scope dessa função, esse canal só envia
//         - Podemos fechar um send channel
// - Exemplo: passando informação de uma função para outra.
package main

import "fmt"

func main() {

    canal := make(chan int)
    go send(canal)
    receive(canal)
    
}


func send(s chan<- int){
    s <- 42
}

func receive(r <-chan int){
    fmt.Println("O numero do canal é", <-r)
}