package main

import (
	"github.com/llRedXD/Go6/db"
	"github.com/llRedXD/Go6/routes"
)

func main() {
	db.ConectDb()
	routes.Handler()
}
