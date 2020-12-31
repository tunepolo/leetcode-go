package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	parameters []string
	expected   string
}

func TestLongestCommonPrefix(t *testing.T) {
	var cases = []Case{
		{
			parameters: []string{"flower", "flow", "flight"},
			expected:   "fl",
		},
		{
			parameters: []string{"dog", "racecar", "car"},
			expected:   "",
		},
		{
			parameters: []string{},
			expected:   "",
		},
		{
			parameters: []string{"flower"},
			expected:   "flower",
		},
		{
			parameters: []string{"flower", "flow", ""},
			expected:   "",
		},
		{
			parameters: []string{"ab", "a"},
			expected:   "a",
		},
	}
	for _, c := range cases {
		result := longestCommonPrefix(c.parameters)
		assert.Equal(t, c.expected, result)
	}
}
