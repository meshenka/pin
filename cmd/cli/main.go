package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/meshenka/pin"
)

func main() {
	iterations := flag.Int("n", 100, "Number of pin code to generate")
	flag.Parse()
	g := pin.NewGenerator()
	var wg sync.WaitGroup
	wg.Add(*iterations)
	for i := 0; i < *iterations; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(g.Generate())
		}()
	}

	wg.Wait()
}
