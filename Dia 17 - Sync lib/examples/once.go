package examples

import (
	"fmt"
	"sync"
)

func dizerOi(currentTime int) {
	fmt.Printf("Oi, essa eh a %vª vez que sou chamado!\n", currentTime)
}

func TestOnce(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("[Testing sync.Once]\n\n")
	fn := sync.Once{}

	fmt.Println("Sem sync.Once")
	for i := 0; i < 10; i++ {
		dizerOi(i + 1)
	}

	fmt.Println("\nCom sync.Once")
	for i := 0; i < 10; i++ {
		fn.Do(func() { dizerOi(i + 1) })
	}

	fmt.Print("\n\n")
}
