package pin

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// GeneratorFunc is a stateless implementation as a function
type GeneratorFunc func() string

func (f GeneratorFunc) Generate() string {
	return f()
}

// NewGenerator creates a Generator service.
func NewGenerator() GeneratorFunc {
	// Generate security rules
	// * cannot be a repeated digit
	// * cannot be a suite of following digits (ascending and descending)
	// * TODO cannot be in the restricted codes
	// So basicaly i randomly pick one number as the first digit and next digit
	// cannot be the same, the previous or the next
	filter := compose(
		nocurrent,
		noprevious,
		nonext,
	)
	return func() string {
		var charset = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		pin := make([]int, 4)
		var current = charset[random.Intn(len(charset))]
		pin[0] = current

		for i := 0; i < 3; i++ {
			// TODO i do not like that, need another way to do that
			charset = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			current = in(filter(charset, current))
			pin[i+1] = current
		}
		return strings.Trim(strings.Join(strings.Split(fmt.Sprint(pin), " "), ""), "[]")
	}
}

func next(v int) int {
	if v == 9 {
		return 0
	}
	return v + 1
}

func previous(v int) int {
	if v == 0 {
		return 9
	}
	return v - 1
}

func in(avail []int) int {
	return avail[random.Intn(len(avail))]
}

func nocurrent(avail []int, current int) []int {
	for i, v := range avail {
		if v == current {
			avail[i] = avail[len(avail)-1]
			return avail[:len(avail)-1]
		}
	}
	return avail
}

func nonext(avail []int, current int) []int {
	n := next(current)
	for i, v := range avail {
		if v == n {
			avail[i] = avail[len(avail)-1]
			return avail[:len(avail)-1]
		}
	}
	return avail
}

func noprevious(avail []int, current int) []int {
	p := previous(current)
	for i, v := range avail {
		if v == p {
			avail[i] = avail[len(avail)-1]
			return avail[:len(avail)-1]
		}
	}
	return avail
}

type filterfunc func([]int, int) []int

func compose(filters ...filterfunc) filterfunc {
	return func(avail []int, current int) []int {
		for _, f := range filters {
			avail = f(avail, current)
		}
		return avail
	}
}
