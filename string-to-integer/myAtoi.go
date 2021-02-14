package myatoi

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var ret int
	var negative bool

	for _, ch := range s {
		switch ch {
		case ' ', '+':
			continue
		case '-':
			negative = true
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
