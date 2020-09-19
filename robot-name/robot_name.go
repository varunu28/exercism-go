package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

var currentNames = map[string]bool{}

// Robot is a factory for generating names
type Robot struct {
	currName string
}

// Name returns name of robot
func (r *Robot) Name() (string, error) {
	if r.currName != "" {
		return r.currName, nil
	}
	s := generateRandomString() + generateRandomNumber()
	for _, ok := currentNames[s]; ok; {
		s = generateRandomString() + generateRandomNumber()
	}
	currentNames[s] = true
	r.currName = s
	return s, nil
}

func generateRandomString() string {
	validLetters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := ""
	for i := 0; i < 2; i++ {
		rand.Seed(time.Now().UnixNano())
		s = s + string(validLetters[rand.Intn(26)])
	}
	return s
}

func generateRandomNumber() string {
	s := ""
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		s = s + strconv.Itoa(randomInt(0, 10))
	}
	return s
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Reset reset robot
func (r *Robot) Reset() {
	r.currName = ""
}
