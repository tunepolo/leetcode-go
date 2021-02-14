package myatoi

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var ret int

	// find first character index
	var index int
FIND_CHAR:
	for i, ch := range []byte(s) {
		switch {
		case ch == ' ':
			continue
		case ch == '+' || ch == '-' || ('0' <= ch && ch <= '9'):
			index = i
			break FIND_CHAR
		default:
			return 0 // invalid input
		}
	}

	for _, ch := range []byte(s[index:]) {
		if ch < '0' || '9' < ch {
			continue
		}

		ch -= '0'
		ret = ret*10 + int(ch)
	}

	if s[index] == '-' {
		ret = -ret
	}

	return ret
}
