package main

import (
	"time"
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := "Alef"
	var version float32 = 1.0

	for {
		showIntroduction(name, version)
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Logs: ")

		case 0:
			fmt.Println("Programa finalizado.")
			os.Exit(0)

		default:
			fmt.Println("Comando inválido.")

		}
	}

}

func showIntroduction(name string, version float32) {
	fmt.Println("Versão do programa: ", version)
	fmt.Println("Olá", name, "qual opção o senhor gostaria de executar?")

	fmt.Println("1- Iniciar o monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- sair do programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento")
	sites := []string{"http://random-status-code.herokuapp.com/", "http://google.com/"}

	for {
		for i, site := range sites {
			fmt.Println("Testando site ", i, ": ", site)
			testingSite(site)
			fmt.Println()

		}
		time.Sleep(5 * time.Second)
	}

}

func testingSite(site string) {

	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("O site está no ar.")
	} else {
		fmt.Println("O site está com problemas. Status code: ", response.StatusCode)
	}
}
