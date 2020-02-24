package main

import (
	"fmt"
	"pin"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const iterations = 1000
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			pin := pin.Generate()
			fmt.Println(pin)
		}()
	}

	wg.Wait()
}
