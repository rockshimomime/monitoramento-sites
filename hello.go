package main

import (
	"fmt"
)

func main() {
	name := "Alef"
	version := 24
	var command int


	fmt.Println("Versão do programa: ", version)
	fmt.Println("Olá", name, "qual opção o senhor gostaria de executar?")

	fmt.Println("1- Iniciar o monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- sair do programa")

	fmt.Scan(&command)

	switch command {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Logs: ")

	case 0:
		fmt.Println("Programa finalizado.")

	default:
		fmt.Println("Comando inválido.")

	}

}
