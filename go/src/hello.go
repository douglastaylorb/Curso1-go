package main

import "fmt"

func main() {
	nome := "Douglas"
	versao := 1.2
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)
}
