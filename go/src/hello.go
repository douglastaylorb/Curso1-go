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

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Até a próxima! ")
	// } else {
	// 	fmt.Println("Não conheço este comando!")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Até a próxima! ")
	default:
		fmt.Println("Não conheço este comando!")
	}
}
