package main

import (
	"fmt"
)

func main() {
	name := "Alef"
	version := 24

	fmt.Println("Versão do programa: ", version)

	fmt.Println("Olá", name, "qual opção o senhor gostaria de executar?")

	fmt.Println("1- Iniciar o monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- sair do programa")

	var command int

	fmt.Scan(&command)

	if command == 1 {
		fmt.Println("Monitorando...")
	} else if command == 2 {
		fmt.Println("Logs: ")
	} else if command == 0 {
		fmt.Println("Programa finalizado.")
	} else {
		fmt.Println("Comando inválido.")
	}

}
