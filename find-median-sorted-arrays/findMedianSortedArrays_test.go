package findMedianSortedArrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	nums1    []int
	nums2    []int
	expected float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	var cases = []Case{
		Case{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		Case{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
	}

	for _, c := range cases {
		result := findMedianSortedArrays(c.nums1, c.nums2)
		assert.Equal(t, c.expected, result)
	}
}
