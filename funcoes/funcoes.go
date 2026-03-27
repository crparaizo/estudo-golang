package main

import "fmt"

func calculosMatematicos(n1, n2 int) (int, int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}

func main() {
	n1 := 10
	n2 := 5
	soma, subtracao, multiplicacao, divisao := calculosMatematicos(n1, n2)
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
}
