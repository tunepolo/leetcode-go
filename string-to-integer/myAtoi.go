package myatoi

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var ret int

	for _, ch := range s {
		ch -= '0'
		ret = ret*10 + int(ch)
	}

	return ret
}
