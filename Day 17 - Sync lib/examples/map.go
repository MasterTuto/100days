package examples

import (
	"fmt"
	"sync"
)

func fill(what string, m *sync.Map, wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 1; i <= 5; i++ {
		m.Store(i, fmt.Sprintf("%s %d", what, i))
	}
	wg.Done()
}

func TestMap(wg *sync.WaitGroup) {
	fmt.Print("[Testing sync.Map]\n\n")

	selfWg := sync.WaitGroup{}

	livingBeings := &sync.Map{}

	go fill("human", livingBeings, &selfWg)
	go fill("plant", livingBeings, &selfWg)
	go fill("cat", livingBeings, &selfWg)
	go fill("monkey", livingBeings, &selfWg)

	selfWg.Wait()

	for i := 1; i <= 5; i++ {
		r, _ := livingBeings.Load(i)

		fmt.Printf("livingBeings[%d]:\n%v\n", i, r)
	}

	fmt.Print("\n\n")

	wg.Done()
}
