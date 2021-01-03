package reverseinteger

// https://leetcode.com/problems/reverse-integer/

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	absX := x
	if x < 0 {
		absX = -x
	}

	result, _ := strconv.Atoi(reverseString(strconv.Itoa(absX)))

	if x > 0 {
		if result > math.MaxInt32 {
			result = 0
		}
	} else {
		result = -result
		if result < math.MinInt32 {
			result = 0
		}
	}
	return result
}

func reverseString(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
