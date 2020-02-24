package main

import (
	"fmt"
	"pin"
)

func main() {
	for i := 0; i < 100; i++ {
		pin := pin.Generate()
		fmt.Println(pin)
	}
}
