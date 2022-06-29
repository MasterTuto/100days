package main

import (
	"fmt"

	"mtbreno.dev/100days.day15/store"
)

func MostrarMenu() {
	fmt.Println("[MENU]")
	fmt.Println("\t1. Ver produtos")
	fmt.Println("\t2. Adicionar produto no carrinho")
	fmt.Println("\t3. Remover produto do carrinho")
	fmt.Println("\t4. Ver carrinho")
	fmt.Println("\t5. Finalizar compra")
	fmt.Println("\n\t0. Sair")
}

func readInt(query string) int {
	var n int

	fmt.Print(query)
	fmt.Scan(&n)

	return n
}

func main() {
	NewProduct := store.NewProductFactory()
	products := store.NewStore()

	products.AddProduct(NewProduct("Coca-Cola", 6.50, nil))
	products.AddProduct(NewProduct("Fanta", 7.50, nil))
	products.AddProduct(NewProduct("Sprite", 2.50, nil))
	products.AddProduct(NewProduct("Guaraná", 7.50, nil))
	products.AddProduct(NewProduct("Água", 2.50, nil))
	products.AddProduct(NewProduct("Cerveja", 8.50, nil))

	cart := store.NewCart()

	MostrarMenu()
	op := readInt("\nDigite a opção: ")

	for op != 0 {
		switch op {
		case 1: // show products
			fmt.Println(products)
		case 2: // add product to cart
			for {
				productId := readInt("Digite o ID do produto: ")
				quantity := readInt("Digite a quantidade: ")

				if _, ok := products.Products[productId]; !ok {
					fmt.Println("Produto não encontrado")
					continue
				}

				cart.AddItem(products.Products[productId], quantity)
				break
			}
		case 3: // remove product from cart
			for {
				productId := readInt("Digite o ID do produto: ")
				quantity := readInt("Digite a quantidade: ")

				if _, ok := products.Products[productId]; !ok {
					fmt.Println("Produto não encontrado")
					continue
				}

				cart.RemoveItem(products.Products[productId], quantity)
				break
			}
		case 4: // show cart
			fmt.Println(cart)
		case 5: // finalize purchase
			fmt.Println("Total: R$", cart.Total())

			cart.Clear()

			fmt.Println("Carrinho limpo")
		}

		MostrarMenu()
		op = readInt("\nDigite a opção: ")
	}
}
