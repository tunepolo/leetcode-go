package romantointeger

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result = 0

	for i := 0; i < len(s); {
		if i+1 == len(s) || m[s[i]] >= m[s[i+1]] {
			result += m[s[i]]
			i++
		} else {
			// Need substraction
			result += m[s[i+1]] - m[s[i]]
			i += 2
		}
	}

	return result
}
