package zigzagconversion

import "strings"

// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	var ret []string

	if numRows == 1 {
		return s
	}

	length := len(s)
	increment := numRows*2 - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j += increment {
			ret = append(ret, string(s[i+j]))

			if i != 0 && i != numRows-1 && j+increment-i < length {
				ret = append(ret, string(s[j+increment-i]))
			}
		}
	}

	return strings.Join(ret, "")
}
