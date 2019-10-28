package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator generate a random number on a range index of given list of prbablity as element.
func Generator(probs []int) (int, error) {
	if probs == nil || len(probs) == 0 {
		return -1, fmt.Errorf("Invalid probability : %v", probs)
	}
	n := len(probs)
	if n == 1 {
		return n, nil
	}
	// construct an sum list from given probabilities
	probSum := make([]int, n)

	// probSum[i] holds sum of all probability[j] for 0 <= j <=i
	probSum[0] = probs[0]
	for i := 1; i < n; i++ {
		probSum[i] = probSum[i-1] + probs[i]
	}

	// generate a random integer from 1 to 100
	// and check where it lies in probSum[]
	r := seedRandom(1, probSum[n-1])

	// based on the comparison result, return corresponding
	// element from the input list

	if r <= probSum[0] { // handle 0'th index separately
		return 0, nil
	}

	for i := 1; i < n; i++ {
		if r > probSum[i-1] && r <= probSum[i] {
			return i, nil
		}
	}

	return -1, nil
}

func seedRandom(min, max int) int {
	// to seed random int properly and unique every time
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min+1) + min
}
