package main

import "fmt"

func main() {
	var peso int = 130
	var altura float32 = 1.82

	fmt.Println("O peso:", peso)
	fmt.Println("A altura:", altura)

	if IMC := float32(peso) / (altura * altura); IMC < 18.5 {
		fmt.Printf("IMC: %v\nResultado: Abaixo do peso\n", IMC)
	} else if IMC >= 18.5 && IMC < 25 {
		fmt.Printf("IMC: %v\nResultado: Peso normal\n", IMC)
	} else if IMC >= 25 && IMC < 30 {
		fmt.Printf("IMC: %v\nResultado: Sobre peso\n", IMC)
	} else if IMC >= 30 && IMC < 35 {
		fmt.Printf("IMC: %v\nResultado: Obesidade Grau I\n", IMC)
	} else if IMC >= 35 && IMC < 40 {
		fmt.Printf("IMC: %v\nResultado: Obesidade Grau II\n", IMC)
	} else if IMC >= 40 {
		fmt.Printf("IMC: %v\nResultado: Obesidade Grau III\n", IMC)
	}

	fmt.Println("IMC nao eh mais acessivel a partir de fora do IF!")

}
