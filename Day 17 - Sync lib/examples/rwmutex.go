package examples

import (
	"fmt"
	"sync"
	"time"
)

func maps(name string, used *sync.RWMutex, value *int, by int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		used.RLock()
		fmt.Printf("Reading... [ %d * %d = %d]\n", (*value), by, (*value)*by)
		used.RUnlock()

		time.Sleep(1 * time.Second)
	}
}

func writeMutex(name string, value *int, used *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 15; i++ {

		used.Lock()
		(*value) = (*value) + 1
		fmt.Printf("Writing... [value = %d]\n", *value)
		used.Unlock()

		time.Sleep(500 * time.Millisecond)
	}
}

func TestRWMutex(wg *sync.WaitGroup) {
	fmt.Println("[Testing sync.RWMutex]")

	defer wg.Done()

	selfWg := sync.WaitGroup{}
	value := 0

	rwMutex := &sync.RWMutex{}

	selfWg.Add(1)
	go maps("reader 1", rwMutex, &value, 5, &selfWg)
	selfWg.Add(1)
	go maps("reader 2", rwMutex, &value, 5, &selfWg)
	selfWg.Add(1)
	go maps("reader 3", rwMutex, &value, 5, &selfWg)

	selfWg.Add(1)
	go writeMutex("writer 1", &value, rwMutex, &selfWg)

	selfWg.Wait()

	fmt.Print("\n\n")
}
