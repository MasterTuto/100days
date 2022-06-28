package main

import "fmt"

func MostrarMenu() {
	Options := []string{"Adicionar", "Subtrair", "Dividir", "Multiplicar", "Sair"}

	fmt.Println("[CALCULADORA 1.0]")
	fmt.Printf("Uso: escolha a opcao e digite o valor\n\n")
	for index, option := range Options {
		fmt.Printf("\t%v. %v\n", (index+1)%len(Options), option)
	}
	fmt.Println()
}

func readInt(query string) int {
	var n int

	fmt.Print(query)
	fmt.Scan(&n)

	return n
}

func readFloat(query string) float64 {
	var n float64

	fmt.Print(query)
	fmt.Scan(&n)

	return n
}

func main() {
	result := readFloat("Digite o valor inicial desejado: ")

	MostrarMenu()
	fmt.Println("RESULTADO ATUAL: ", result)
	op := readInt("Digite uma opcao valida: ")
	for op != 0 {
		switch op {
		case 1, 2, 3, 4:
			value := readFloat("Digite um valor: ")
			result = Operations[op](result, value)
		default:
			fmt.Println("Erro")
		}

		MostrarMenu()
		fmt.Println("RESULTADO ATUAL: ", result)
		op = readInt("Digite uma opcao valida: ")
	}

	fmt.Println("Obrigado por usar!")
}
