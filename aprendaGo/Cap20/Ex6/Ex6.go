// - Crie um programa que demonstra seu OS e ARCH.
// - Rode-o com os seguintes comandos:
//   - go run X
//   - go build
//   - go install
package main

import (
	"fmt"
	"runtime"
)

func main() {

    fmt.Println("OS:",runtime.GOOS)
    fmt.Println("ARCH:",runtime.GOARCH)
}