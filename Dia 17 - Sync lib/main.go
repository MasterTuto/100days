package main

import (
	"sync"

	"mtbreno.dev/100days.day17/examples"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	examples.TestCond(&wg)
	wg.Add(1)
	examples.TestMap(&wg)
	wg.Add(1)
	examples.TestMutex(&wg)
	wg.Add(1)
	examples.TestOnce(&wg)
	wg.Add(1)
	examples.TestPool(&wg)
	wg.Add(1)
	examples.TestRWMutex(&wg)

	wg.Wait()
}
