package myatoi

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var ret int
	var negative bool

	for _, ch := range []byte(s) {
		switch {
		case ch == '-':
			negative = true
			continue
		case ch < '0' || '9' < ch:
			continue
		}

		ch -= '0'
		ret = ret*10 + int(ch)
	}

	if negative {
		ret = -ret
	}

	return ret
}
