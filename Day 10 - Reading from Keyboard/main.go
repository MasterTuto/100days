package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	name   string
	age    int
	birth  time.Time
	height float32
	weight float32
}

func (p Person) GetBMI() float32 {
	return p.weight / (p.height * p.height)
}

func (p Person) HowFatAmI() string {
	bmi := p.GetBMI()

	switch {
	case bmi < 18.5:
		return "magreza"
	case bmi >= 18.5 && bmi < 25:
		return "normal"
	case bmi >= 25 && bmi < 30:
		return "sobrepeso"
	case bmi >= 30 && bmi < 40:
		return "obesidade"
	case bmi >= 40:
		return "obesidade grave"
	}

	return "unknown"
}

func ReadPerson() Person {
	var (
		name             string
		age              int
		year, month, day int
		height, weight   float32
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite seu nome: ")
	name, _ = reader.ReadString('\n')

	for {
		fmt.Print("Digite sua idade: ")
		temp, _ := reader.ReadString('\n')

		var err error

		age, err = strconv.Atoi(strings.TrimSpace(temp))

		if err != nil {
			fmt.Print(err)
			continue

		}

		break
	}

	for {
		fmt.Print("Digite sua data de nascimento (e.g. DD/MM/AAAA): ")
		temp, _ := reader.ReadString('\n')

		birthDate := strings.Split(temp, "/")

		var err1, err2, err3 error
		day, err1 = strconv.Atoi(strings.TrimSpace(birthDate[0]))
		month, err2 = strconv.Atoi(strings.TrimSpace(birthDate[1]))
		year, err3 = strconv.Atoi(strings.TrimSpace(birthDate[2]))

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Entrada invalida, tente novamente.")
			continue
		}

		break
	}

	for {
		fmt.Print("Digite sua altura: ")
		temp, _ := reader.ReadString('\n')
		temp2, err := strconv.ParseFloat(strings.TrimSpace(temp), 32)

		if err != nil {
			continue
		}

		height = float32(temp2)
		break
	}

	for {
		fmt.Print("Digite seu peso: ")
		temp, _ := reader.ReadString('\n')
		temp2, err := strconv.ParseFloat(strings.TrimSpace(temp), 32)

		if err != nil {
			continue
		}

		weight = float32(temp2)
		break
	}

	return Person{
		name:   name,
		age:    age,
		birth:  time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
		height: height,
		weight: weight,
	}
}

func (p Person) String() string {
	return fmt.Sprintf(
		"[PESSOA]\nNome: %v\nIdade: %v, Data de Nascimento: %v\nAltura: %v\nPeso: %v\nIMC: %v, pelo seu IMC, vc tem %v.\n",
		p.name,
		p.age,
		p.birth,
		p.height,
		p.weight,
		p.GetBMI(),
		p.HowFatAmI(),
	)
}

func main() {

	fmt.Print("Bem vindo ao seu cadastro de gordura!\nSe apresente:\n\n")
	person := ReadPerson()

	fmt.Print("Analisando...")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(".")

	fmt.Printf("\nConclusao de seus dados:\n%v\n", person)

}
