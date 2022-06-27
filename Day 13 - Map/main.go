package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf("#%v - %v (Cost: R$ %v)", p.Id, p.Name, p.Price)
}

type Order struct {
	Items map[Product]int
}

func NewProductFactory() func(string, float64) Product {
	lastIndex := 0

	return func(Name string, Price float64) Product {
		lastIndex += 1
		return Product{
			Id:    lastIndex,
			Name:  Name,
			Price: Price,
		}
	}
}

func NewOrder() *Order {
	return &Order{
		Items: make(map[Product]int),
	}
}

func (order *Order) AddItem(p Product, quantity int) {
	_, found := order.Items[p]

	if found {
		order.Items[p]++
	} else {
		order.Items[p] = 1
	}
}

func (order *Order) RemoveItem(p Product) bool {
	_, found := order.Items[p]

	delete(order.Items, p)

	return found
}

func (order Order) Total() float64 {
	result := 0.0
	for product, quantity := range order.Items {
		result += product.Price * float64(quantity)
	}

	return result
}

func main() {
	NewProduct := NewProductFactory()

	rice := NewProduct("Rice", 5.80)
	bean := NewProduct("Bean", 8.50)
	pasta := NewProduct("Pasta", 6.60)

	order := NewOrder()

	fmt.Printf("Adding 3 units of %v...\n", rice)
	order.AddItem(rice, 3)
	fmt.Printf("Adding 5 units of %v...\n", bean)
	order.AddItem(bean, 5)
	fmt.Printf("Adding 2 units of %v...\n", pasta)
	order.AddItem(pasta, 2)

	fmt.Printf("\nOrder total: %v\n", order.Total())

	fmt.Printf("\nRemoving bean...\n")
	order.RemoveItem(bean)

	fmt.Printf("\nOrder total... again: %v\n", order.Total())

}
