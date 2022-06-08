package main

import (
	"fmt"
)

type Date struct {
	Day, Month, Year int
}

func (d Date) String() string {
	return fmt.Sprintf("%v/%v/%v", d.Day, d.Month, d.Year)
}

type Animal struct {
	Name         string
	AmountWalked int64
	Birth        Date
}

func (a Animal) String() string {
	return fmt.Sprintf(
		"Name: %v\nAmount Walked: %v mts\nBirth: %v",
		a.Name,
		a.AmountWalked,
		a.Birth,
	)
}

type Dog struct {
	Animal
	FurColor string
}

type Cat struct {
	Animal
	FurColor   string
	HasStripes bool
}

type iAnimal interface {
	GetName() string
	GetAmountWalked() int64
	MakeSound() string
	Walk()
	WalkWithoutReference()
	String() string
}

/* Dog methods */
func (dog Dog) GetName() string {
	return dog.Name
}

func (dog Dog) MakeSound() string {
	return "Woof Woof"
}

func (dog Dog) GetAmountWalked() int64 {
	return dog.AmountWalked
}

func (dog *Dog) Walk() {
	dog.AmountWalked += 3
}

func (dog Dog) WalkWithoutReference() {
	dog.AmountWalked += 3
}

func (dog Dog) String() string {
	return fmt.Sprintf(
		"[Dog Animal]\n%v\nFurColor: %v",
		dog.Animal,
		dog.FurColor,
	)
}

/* Cat methods */
func (cat Cat) GetName() string {
	return cat.Name
}

func (cat Cat) MakeSound() string {
	return "Meow"
}

func (cat Cat) GetAmountWalked() int64 {
	return cat.AmountWalked
}

func (cat *Cat) Walk() {
	cat.AmountWalked += 1
}

func (cat Cat) WalkWithoutReference() {
	cat.AmountWalked += 1
}

func (cat Cat) String() string {
	return fmt.Sprintf(
		"[Cat Animal]\n%v\nFurColor: %v\nHas Stripe: %v",
		cat.Animal,
		cat.FurColor,
		cat.HasStripes,
	)
}

func CallAnimal(animal iAnimal) {
	var distance uint = 13

	fmt.Println("Person: Hey buddy!")
	fmt.Printf("%v: %v\n", animal.GetName(), animal.MakeSound())

	for distance > 0 {
		distance--
		animal.Walk()
	}

	fmt.Printf("Amount walked: %v meters\n", animal.GetAmountWalked())
}

func NewDog(Name, FurColor string, Day, Month, Year int) Dog {
	return Dog{
		Animal{
			Name:         Name,
			AmountWalked: 0,
			Birth: Date{
				Day,
				Month,
				Year,
			},
		},
		FurColor,
	}
}

func NewCat(Name, FurColor string, Day, Month, Year int, HasStripes bool) Cat {
	return Cat{
		Animal{
			Name:         Name,
			AmountWalked: 0,
			Birth: Date{
				Day,
				Month,
				Year,
			},
		},
		FurColor,
		HasStripes,
	}
}

func main() {
	beethoven := NewDog("Beethoven", "Orange", 10, 12, 2019)
	xanax := NewCat("Xanax", "Black", 5, 10, 2021, true)

	fmt.Println("Calling Beethoven...")
	CallAnimal(&beethoven)
	fmt.Println("\nCalling Xanax...")
	CallAnimal(&xanax)

	fmt.Println("\n\nWalking without reference example (Beethoven): ")
	fmt.Println("Amount Walked before:", beethoven.AmountWalked)
	beethoven.WalkWithoutReference()
	fmt.Print("Amount Walked after:", beethoven.AmountWalked, "\n")

	sherlock := &xanax
	sherlock.Name = "Sherlock"
	fmt.Println("\n\nSherlock Name:", sherlock.Name)
	fmt.Println("Xanax Name:", xanax.Name)

	fmt.Print("\n====== All Animals ======\n\n")

	fmt.Printf("%v\n\n%v\n\n%v", beethoven, xanax, sherlock)
}
