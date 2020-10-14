package leetcode

import "testing"

var testCases = []struct {
	name string
	n    int
	pick int
}{
	{"picked 6 from 1-10", 10, 6},
}

func guess(guess, pick int) int {
	if pick < guess {
		return -1
	}
	if pick > guess {
		return 1
	}
	return 0
}

func TestGuessNumber(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := guessNumber(tc.n, tc.pick)
			if actual != tc.pick {
				t.Errorf("Expected: %d, but got %d", tc.pick, actual)
			}
		})
	}
}
