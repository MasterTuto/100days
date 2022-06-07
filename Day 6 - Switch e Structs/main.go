package main

import "fmt"

type Address struct {
	Number, Street, Province, State, Country string
}

type PersonType int64

type Person struct {
	Name           string
	Age            int
	Height, Weight float32
}

type PhysicalPerson struct {
	Cpf string
	Person
}

type JuridicPerson struct {
	Cnpj string
	Person
}

const (
	Physical PersonType = iota
	Juridic
)

func New_PhysicalPerson(
	Name string, Age int, Height, Weight float32,
	Cpf string) PhysicalPerson {

	Person := Person{Name, Age, Height, Weight}

	return PhysicalPerson{Cpf: Cpf, Person: Person}
}

func New_JuridicPerson(
	Name string, Age int, Height, Weight float32,
	Cnpj string) JuridicPerson {

	Person := Person{Name, Age, Height, Weight}

	return JuridicPerson{Cnpj: Cnpj, Person: Person}
}

func main() {
	Breno := New_PhysicalPerson("Breno", 21, 1.73, 80, "000.001.002-03")
	Carlos := New_JuridicPerson("Carlos", 18, 1.83, 95, "12.1241.241/0000-00")

	switch Breno.Name {
	case "Gustavo":
		fmt.Println("O nome de Breno eh Gustavo")
	case "Breno":
		fmt.Println("O nome de Breno eh Breno")
	default:
		fmt.Println("O nome de Breno eu nao sei")
	}

	switch Name := Carlos.Name; Name {
	case "Gustavo":
		fmt.Println("O nome de Carlos eh Gustavo")
	case "Carlos":
		fmt.Println("O nome de Carlos eh Carlos")
	default:
		fmt.Println("O nome de Carlos eu nao sei")
	}

	switch Age := Carlos.Age; {
	case Age < 18:
		fmt.Println("Carlos tem menos de 18 anos")
	case Age == 18:
		fmt.Println("Carlos ja pode ser preso")
		fallthrough
	case Age > 18:
		fmt.Println("Carlos tem mais de 18 anos!")
	}

Loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			break // skips even numbers
		case i == 7:
			break Loop // break that for loop
		default:
			fmt.Println("i:", i)
		}
	}

}
