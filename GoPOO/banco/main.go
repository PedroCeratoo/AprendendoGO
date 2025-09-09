package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	ContaPedro := ContaCorrente{"Pedro", 589, 123456, 125.50}
	fmt.Println(ContaPedro)

	var contaDaLara *ContaCorrente
	contaDaLara = new(ContaCorrente)
	contaDaLara.titular = "Lara"

	fmt.Println(*contaDaLara)

}
