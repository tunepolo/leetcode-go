package palindromenumber

// https://leetcode.com/problems/palindrome-number/

import (
	"strconv"
)

func isPalindrome(x int) bool {
	palindromeNumber, _ := strconv.Atoi(reverseString(strconv.Itoa(x)))
	return x == palindromeNumber
}

func reverseString(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	reverseX := 0
	for n := x; n != 0; n /= 10 {
		reverseX = reverseX*10 + n%10
	}

	return x == reverseX
}
