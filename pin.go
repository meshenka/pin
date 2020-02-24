package pin

import (
	"math/rand"
	"strings"
	"time"
)

// Generator expose the service API
type Generator interface {
	Get() string
}

// Generate security rules
// * cannot be a repeated digit
// * cannot be a suite of following digits (ascending and descending)
// * cannot be in the restricted codes
// So basicaly i randomly pick one number as the first digit and next digit
// cannot be the same, the previous or the next
func Generate() string {

	var pin []string
	var current = first()

	pin = append(pin, current)

	current = chooseNext(current)
	pin = append(pin, current)

	current = chooseNext(current)
	pin = append(pin, current)

	current = chooseNext(current)
	pin = append(pin, current)

	return strings.Join(pin, "")
}

func next(s string) string {
	nextMap := make(map[string]string)

	nextMap["0"] = "1"
	nextMap["1"] = "2"
	nextMap["2"] = "3"
	nextMap["3"] = "4"
	nextMap["4"] = "5"
	nextMap["5"] = "6"
	nextMap["6"] = "7"
	nextMap["7"] = "8"
	nextMap["8"] = "9"
	nextMap["9"] = "0"

	return nextMap[s]
}

func previous(s string) string {
	nextMap := make(map[string]string)

	nextMap["0"] = "9"
	nextMap["1"] = "0"
	nextMap["2"] = "1"
	nextMap["3"] = "2"
	nextMap["4"] = "3"
	nextMap["5"] = "4"
	nextMap["6"] = "5"
	nextMap["7"] = "6"
	nextMap["8"] = "7"
	nextMap["9"] = "8"

	return nextMap[s]
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func inCharset(charset string) string {
	b := make([]byte, 1)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func first() string {
	return inCharset("1234567890")
}

func chooseNext(current string) string {
	charset := "1234567890"
	//remove current next and previous
	charset = strings.Replace(charset, current, "", -1)
	charset = strings.Replace(charset, next(current), "", -1)
	charset = strings.Replace(charset, previous(current), "", -1)
	return inCharset(charset)
}
