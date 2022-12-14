package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 2

func main() {
	Introducao()

	for {
		Menu()
		comando := Opcoes()

		switch comando {
		case 1:
			Monitoramento()
		case 2:
			MostraLog()
		case 0:
			fmt.Println("Saindo")
			os.Exit(0)
		default:
			fmt.Println("Opção Invalida")
			os.Exit(-1)
		}

		fmt.Println("Valor Recebido", comando)
	}

}

func Menu() {
	fmt.Println("1- Inciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func Introducao() {
	nome := "Gabriel"
	versao := 2.1
	fmt.Printf("Bem vindo %v\nO sistema esta na versão %v\n", nome, versao)
}

func Opcoes() int {
	comando := 0
	fmt.Scan(&comando)
	return comando
}

func Monitoramento() {
	fmt.Println("Monitorando....")
	sites := LerNomeDosSites()
	for i := 0; i < delay; i++ {
		ConsultaSites(sites)
		time.Sleep(2 * time.Second)
	}

}

func ConsultaSites(sites []string) {
	fmt.Println("")
	for _, v := range sites {
		r, err := http.Get(v)
		if err != nil {
			log.Println("Erro:", err)
		}
		if r.StatusCode == 200 {
			fmt.Println("Site:", v, "Foi carregado com sucesso. Status:", r.StatusCode)
			CriaLog(v, r.StatusCode, true)
		} else {
			fmt.Println("Site:", v, "Está com problemas. Status:", r.StatusCode)
		}
	}
}

func LerNomeDosSites() []string {
	site := []string{}

	v, err := os.Open("sites.txt")
	if err != nil {
		log.Println("Erro:", err)
	}
	defer v.Close()

	teste := bufio.NewReader(v)
	for {
		s, err := teste.ReadString('\n')
		if err != nil {
			break
		}
		s = strings.TrimSpace(s)
		site = append(site, string(s))
	}
	return site
}

func CriaLog(site string, status int, conexao bool) {
	a, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Erro:", err)
	}

	tempo := time.Now().Format("02/01/2006 15:04:05")

	a.WriteString(tempo + " Site: " + site + " Code: " + strconv.FormatInt(int64(status), 10) + " Conexão: " + strconv.FormatBool(conexao) + "\n")

}


func MostraLog()  {
	fmt.Println("Logs")
	v, err := os.ReadFile("log.txt")
	if err != nil {
		log.Println("Erro:", err)
	}
	fmt.Println(string(v))
}