package myatoi

import "math"

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var ret int

	s = filterInteger(s)
	if len(s) == 0 {
		return 0
	}

	digitStart := 0
	if s[0] == '+' || s[0] == '-' {
		digitStart = 1
	}
	for _, ch := range []byte(s[digitStart:]) {
		ch -= '0'
		ret = ret*10 + int(ch)
		if ret > math.MaxInt32 {
			break
		}
	}

	if s[0] == '-' {
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
	// find first charcter
	start := 0
	for ; start < len(s); start++ {
		ch := s[start]
		if ch == ' ' {
			continue
		} else if ch == '+' || ch == '-' {
			break
		} else if '0' <= ch && ch <= '9' {
			break
		} else {
			return ""
		}
	}

	// find end character
	end := start + 1
	for ; end < len(s); end++ {
		ch := s[end]
		if ch < '0' || '9' < ch {
			break
		}
	}

	// digits not found after +/-
	s = s[start:end]
	if s == "+" || s == "-" {
		return ""
	}

	return s
}
