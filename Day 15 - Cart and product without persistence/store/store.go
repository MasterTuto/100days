package store

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

type Product struct {
	id               int
	name             string
	price            float64
	discountFunction func(int) float64
}

type CartItem struct {
	product  Product
	quantity int
}

func (ci CartItem) SubTotal() float64 {
	result := ci.product.price * float64(ci.quantity)

	if ci.product.discountFunction != nil {
		return ci.product.discountFunction(ci.quantity) * result
	}

	return result
}

type Cart struct {
	items map[int]*CartItem
}

func (c *Cart) AddItem(p Product, n int) {
	if item, found := c.items[p.id]; found {
		item.quantity += n
	} else {
		c.items[p.id] = &CartItem{product: p, quantity: n}
	}
}

func (c *Cart) RemoveItem(p Product, n int) error {
	item, found := c.items[p.id]
	if found {
		item.quantity -= min(n, item.quantity)
	} else {
		return fmt.Errorf("Produto %d não encontrado no carrinho", p.id)
	}

	if item.quantity == 0 {
		delete(c.items, p.id)
	}

	return nil
}

func (c Cart) Total() float64 {
	total := 0.0

	for _, item := range c.items {
		total += item.SubTotal()
	}

	return total
}

func (c Cart) String() string {
	result := ""

	for _, item := range c.items {
		result += fmt.Sprintf("%d - %s (Qntd: %d)- R$ %.2f\n", item.product.id, item.product.name, item.quantity, item.SubTotal())
	}

	// append total
	result += fmt.Sprintf("Total: R$ %.2f\n", c.Total())

	return result
}

func (c Cart) Clear() {
	c.items = make(map[int]*CartItem)
}

type Store struct {
	Products map[int]Product
}

func NewStore() Store {
	return Store{Products: make(map[int]Product)}
}

func (s Store) AddProduct(p Product) {
	s.Products[p.id] = p
}

func (s Store) String() string {
	result := ""

	for _, p := range s.Products {
		result += fmt.Sprintf("%d - %s - R$ %.2f\n", p.id, p.name, p.price)
	}

	return result
}

func NewCart() Cart {
	return Cart{items: make(map[int]*CartItem)}
}

func NewProductFactory() func(string, float64, func(int) float64) Product {
	currentId := 1

	return func(s string, f float64, df func(int) float64) Product {
		df = (func() func(int) float64 {
			if df == nil {
				return func(n int) float64 { return 1 }
			}

			return df
		})()

		defer func() { currentId++ }()

		return Product{
			id:               currentId,
			name:             s,
			price:            f,
			discountFunction: df,
		}
	}
}
