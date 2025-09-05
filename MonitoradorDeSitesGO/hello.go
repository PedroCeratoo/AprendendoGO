package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs!")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não disponível !")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Pedro"
	versao := 1.1
	fmt.Println("Olá. Sr. ", nome)
	fmt.Println("Esse programa está na versão: ", versao)
	fmt.Println()
}

func lerComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("\nO comando lido foi:", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("\n1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("MONITORANDO..!")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {

		fmt.Println("\n Carregando monitoramento !")
		for i, site := range sites {
			fmt.Println("Posição:", i, " Site:", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está passando por problemas, tente novamente mais tarde!")
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("Sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites
}
