package core

import (
	"fmt"

	"github.com/llRedXD/factoryDesign/database"
	"github.com/llRedXD/factoryDesign/webserver"
)


type core struct {
}

func (c *core) Start() {
	fmt.Println("> [core] Starting...")
	database.New().Start()
	webserver.New().Start()
	fmt.Println("> [core] Starting Done! System running!")
}

func (c *core) Stop() {
	fmt.Println("> [core] Stopping...")
	fmt.Println("> [core] Stopping Done! System stopped!")
}

func New() *core {
	return &core{}
}