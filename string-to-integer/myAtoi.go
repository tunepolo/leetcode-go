package myatoi

import (
	"bytes"
	"math"
	"strings"
)

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	s = filterInteger(s)
	if len(s) == 0 {
		return 0
	}

	var negative bool
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	var ret int
	for _, ch := range []byte(s) {
		ch -= '0'
		ret = ret*10 + int(ch)
		if ret > math.MaxInt32 {
			break
		}
	}

	if negative {
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

func filterInteger(s string) string {
	// check the first character is +/- or digit
	s = strings.Trim(s, " ")
	if len(s) == 0 || !bytes.ContainsAny([]byte("+-0123456789"), string(s[0])) {
		return ""
	}

	// find end position
	end := 1
	for ; end < len(s); end++ {
		ch := s[end]
		if ch < '0' || '9' < ch {
			break
		}
	}

	return strings.Trim(s[:end], "+")
}
