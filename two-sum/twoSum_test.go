package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	var cases = []Case{
		Case{nums: []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
	}

	for _, c := range cases {
		result := twoSum(c.nums, c.target)
		assert.Equal(t, c.expected, result)
	}
}
