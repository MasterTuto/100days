package examples

import (
	"fmt"
	"sync"
)

func addWithoutMutex(x *int, n int, done *sync.WaitGroup) {
	defer done.Done()

	for i := 0; i < 100000; i++ {
		(*x) += n
	}

}

func addWithMutex(x *int, n int, m *sync.Mutex, done *sync.WaitGroup) {
	defer done.Done()

	for i := 0; i < 100000; i++ {
		m.Lock()
		(*x) += n
		m.Unlock()
	}

}

func TestMutex(wg *sync.WaitGroup) {

	done := sync.WaitGroup{}
	mutex := sync.Mutex{}

	fmt.Print("[Testing sync.Mutex]\n\n")
	fmt.Println("Operations (addition and subtraction below are parallell):")
	fmt.Println("x = 0 (initial)")
	fmt.Println("x = x + 1 + 1 + 1 + ... 1 (n times)")
	fmt.Println("x = x - 1 - 1 - 1 - ... 1 (n times)")

	fmt.Println("\nWithout mutex (n = 100000)...")
	x := 0
	done.Add(1)
	go addWithoutMutex(&x, 1, &done)
	done.Add(1)
	go addWithoutMutex(&x, -1, &done)
	done.Wait()
	fmt.Printf("After, x = %v\n", x)

	y := 0
	fmt.Println("\nWith mutex (n = 100000)...")
	done.Add(1)
	go addWithMutex(&y, 1, &mutex, &done)
	done.Add(1)
	go addWithMutex(&y, -1, &mutex, &done)
	done.Wait()
	fmt.Printf("After, y = %v\n", y)

	fmt.Print("\n\n")
	wg.Done()
}
