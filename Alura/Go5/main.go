package main

import (
	"github.com/llRedXD/go5/db"
	"github.com/llRedXD/go5/routes"
)

func main() {
	db.ConectDb()
	routes.Handler()
}
