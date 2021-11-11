package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	for {
		var opcao int

		menu()

		opcao = leOpcao()

		switch opcao {
		case 1:
			dado()
		case 0:
			os.Exit(0)
		}
	}
}

func dado() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 7

	dado := rand.Intn(max-min) + min

	fmt.Println("Dado:", dado)
}

func menu() {
	fmt.Println("1 - Jogar Dado")
	fmt.Println("0 - Sair")
}

func leOpcao() int {
	var opcao int
	fmt.Scan(&opcao)
	fmt.Println("")

	return opcao
}
