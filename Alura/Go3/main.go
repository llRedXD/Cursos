package main

import (
	"net/http"

	"github.com/llRedXD/Go3/routes"
)


func main() {
	routes.CarregandoRotas()
	http.ListenAndServe(":8000", nil)
}

