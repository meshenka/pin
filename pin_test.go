package pin_test

import (
	"testing"

	"github.com/meshenka/pin"
	"github.com/stretchr/testify/assert"
)

func BenchmarkGenerator(b *testing.B) {
	SUT := pin.NewGenerator()
	for i := 0; i < b.N; i++ {
		SUT.Generate()
	}
}

func TestGenerator(t *testing.T) {

	SUT := pin.NewGenerator()
	t.Run("pin is length 4", func(t *testing.T) {
		pin := SUT.Generate()
		assert.Len(t, pin, 4)
	})

	t.Run("pin is number only", func(t *testing.T) {
		pin := SUT.Generate()
		assert.Regexp(t, "^[0-9]{4}$", pin)
	})

	// t.Run("no duplicate number", func(t *testing.T) {
	// 	for i := 0; i < 10; i++ {
	// 		pin := SUT.Generate()
	// 	}
	// })
}
