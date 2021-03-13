package pin

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	nextMap = map[int]int{
		0: 1,
		1: 2,
		2: 3,
		3: 4,
		4: 5,
		6: 7,
		7: 8,
		8: 9,
		9: 0,
	}
	previousMap = map[int]int{
		0: 9,
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
		6: 5,
		7: 6,
		8: 7,
		9: 8,
	}
)

// Generate security rules
// * cannot be a repeated digit
// * cannot be a suite of following digits (ascending and descending)
// * cannot be in the restricted codes
// So basicaly i randomly pick one number as the first digit and next digit
// cannot be the same, the previous or the next
func Generate() string {

	var current = first()
	pin := []int{current}

	for i := 0; i < 3; i++ {
		current = chooseNext(current)
		pin = append(pin, current)
	}

	//return strings.Join(pin, "")
	return fmt.Sprintf("%v", pin)
}

func next(s int) int {
	return nextMap[s]
}

func previous(s int) int {
	return previousMap[s]
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func inCharset(cs []int) int {
	return cs[seededRand.Intn(len(cs))]
}

func first() int {
	return inCharset(charset)
}

var charset = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func chooseNext(current int) int {

	_charset := []int{}
	next := next(current)
	prev := previous(current)
	for v := range charset {
		if v != current && v != next && v != prev {
			_charset = append(_charset, v)
		}
	}

	return inCharset(_charset)
}
