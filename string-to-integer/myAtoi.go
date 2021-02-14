package myatoi

import "math"

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var ret int
	var firstByte byte

	// find first character
	for i, ch := range []byte(s) {
		if ch == ' ' {
			continue
		} else if ch == '+' || ch == '-' {
			firstByte = ch
			s = s[i+1:]
			break
		} else if '0' <= ch && ch <= '9' {
			s = s[i:]
			break
		} else {
			return 0 // invalid input
		}
	}

	for _, ch := range []byte(s) {
		if ch < '0' || '9' < ch {
			break
		}

		ch -= '0'
		ret = ret*10 + int(ch)
		if ret > math.MaxInt32 {
			break
		}
	}

	if firstByte == '-' {
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
