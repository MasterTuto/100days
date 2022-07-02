package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func putRandomString(p *sync.Pool, wg *sync.WaitGroup) {
	defer wg.Done()
	randomString := ""

	for i := 0; i < 10; i++ {
		for i := 0; i < 15; i++ {
			initialChar := rand.Intn(2)*(97-65) + 65
			randomString += string(rand.Intn(26) + initialChar)
		}

		p.Put(fmt.Sprintf("%s - Generated word n. %d", randomString, i+1))
		time.Sleep(500 * time.Millisecond)

		randomString = ""
	}
}

func getRandomString(p *sync.Pool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		str := p.Get()

		fmt.Printf("[%d] Read new value: %v\n", i+1, str)
	}
}

func TestPool(wg *sync.WaitGroup) {
	fmt.Println("[Testing sync.Pool]")

	pool := sync.Pool{
		New: func() any {
			return "Palavra nao gerada"
		},
	}

	selfWg := sync.WaitGroup{}

	selfWg.Add(1)
	go putRandomString(&pool, &selfWg)
	selfWg.Add(1)
	go getRandomString(&pool, &selfWg)

	selfWg.Wait()

	wg.Done()
	fmt.Print("\n\n")
}
