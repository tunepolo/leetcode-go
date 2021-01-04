package longestpalindromicsubstring

// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	start := 0
	end := 1

	for i := 0; i < len(s); i++ {
		length := findPalindromeFromMiddle(s, i, i) // Odd padrindrome case : XaX
		l := findPalindromeFromMiddle(s, i, i+1)    // Even parindrome case : XaaX
		if length < l {
			length = l
		}

		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2 + 1
		}
	}

	return s[start:end]
}

// return length of found longest palindrome
func findPalindromeFromMiddle(s string, left, right int) int {
	for 0 <= left && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
