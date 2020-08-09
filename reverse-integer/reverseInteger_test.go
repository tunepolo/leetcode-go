package reverseinteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	parameters int
	expected   int
}

func TestReverse(t *testing.T) {
	cases := []Case{
		Case{
			parameters: 123,
			expected:   321,
		},
		Case{
			parameters: -123,
			expected:   -321,
		},
		Case{
			parameters: 120,
			expected:   21,
		},
		Case{
			parameters: 1534236469,
			expected:   0,
		},
		Case{
			parameters: -1147483647,
			expected:   0,
		},
	}

	for _, c := range cases {
		result := reverse(c.parameters)
		assert.Equal(t, c.expected, result)
	}
}
