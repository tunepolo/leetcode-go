package myatoi

import "math"

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var ret int

	// find first character index
	var index int
	for i, ch := range []byte(s) {
		if ch == ' ' {
			continue
		} else if ch == '+' || ch == '-' || ('0' <= ch && ch <= '9') {
			index = i
			break
		} else {
			return 0 // invalid input
		}
	}

	for _, ch := range []byte(s[index:]) {
		if ch < '0' || '9' < ch {
			if ch == '+' || ch == '-' {
				continue
			}

			break
		}

		ch -= '0'
		ret = ret*10 + int(ch)
	}

	if s[index] == '-' {
		ret = -ret
	}

	// range adjustment
	if ret > math.MaxInt32 {
		ret = math.MaxInt32
	} else if ret < math.MinInt32 {
		ret = math.MinInt32
	}

	return ret
}
