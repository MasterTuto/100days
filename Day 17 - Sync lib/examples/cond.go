package examples

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func read(name string, cond *sync.Cond) {
	log.Printf("%v trying to read...\n", name)

	cond.L.Lock()

	cond.Wait()

	cond.L.Unlock()
	log.Printf("%v began reading...\n", name)

	time.Sleep(1 * time.Second)

	log.Printf("%v finished reading...\n", name)
}

func write(name string, cond *sync.Cond) {
	log.Printf("%v started writing\n", name)

	cond.L.Lock()

	cond.L.Unlock()

	time.Sleep(5 * time.Second)

	log.Printf("%v wakes all\n", name)
	cond.Broadcast()
	log.Printf("%v finished writing\n", name)

}

func TestCond(wg *sync.WaitGroup) {
	fmt.Print("[Testing sync.Cond]\n\n")

	cond := sync.NewCond(&sync.Mutex{})

	go read("reader 1", cond)
	go read("reader 2", cond)
	go read("reader 3", cond)
	go read("reader 4", cond)

	write("writer", cond)

	time.Sleep(7 * time.Second)

	fmt.Print("\n\n")
	wg.Done()
}
