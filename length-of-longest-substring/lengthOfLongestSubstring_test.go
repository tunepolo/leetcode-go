package lengthoflongestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	str      string
	expected int
}

func TestTwoSum(t *testing.T) {
	var cases = []Case{
		Case{
			str:      "abcabcbb",
			expected: 3,
		},
		Case{
			str:      "bbbbb",
			expected: 1,
		},
		Case{
			str:      "pwwkew",
			expected: 3,
		},
	}

	for _, c := range cases {
		result := lengthOfLongestSubstring(c.str)
		assert.Equal(t, c.expected, result)
	}
}
