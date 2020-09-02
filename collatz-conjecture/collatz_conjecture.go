package collatzconjecture

import "errors"

// CollatzConjecture returns number of steps to reach 0
func CollatzConjecture(input int) (int, error) {
	if input < 0 {
		return -1, errors.New("negative value is an error")
	}
	if input == 0 {
		return -1, errors.New("zero is an error")
	}
	numOfSteps := 0
	for input > 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}
		numOfSteps = numOfSteps + 1
	}
	return numOfSteps, nil
}
