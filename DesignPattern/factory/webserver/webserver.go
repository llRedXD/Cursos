package webserver

import (
	"fmt"
)

type webserver struct {
}

func (w *webserver) Start() {
	fmt.Println("> [webserver] Starting...")
	fmt.Println("> [webserver] Starting Done! Webserver running!")
}

func (w *webserver) Stop() {
	fmt.Println("> [webserver] Stopping...")
	fmt.Println("> [webserver] Stopping Done! Webserver stopped!")
}

func New() *webserver {
	return &webserver{}
}