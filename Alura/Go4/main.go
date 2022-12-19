package main

import (
	"fmt"

	"github.com/llRedXD/aluraRestApi/routes"
)

func main() {
	fmt.Println("Inciando Servidor Rest")
	routes.HandleRequest()
}
