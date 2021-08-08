package impl

import (
	"strings"
)

// https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	return strStr_BM(haystack, needle)
}

// Boyer–Moore
// https://tech.pjin.jp/blog/2020/11/13/stringsearch-2/
func strStr_BM(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}

	table := buidSkiptable(needle)

	i := 0
loop:
	for i+needleLen <= haystackLen {
		for j := needleLen - 1; j >= 0; j-- {
			if haystack[i+j] != needle[j] {
				skip := 0
				if _, ok := table[haystack[i+j]]; !ok {
					// haystackに無い文字で不一致となった場合
					if j == needleLen-1 {
						// 末尾で見つかった場合はneedleの文字数分ずらす
						skip = needleLen
					} else {
						// 途中で見つかった場合はテキストの注目点を右にずらす
						skip = j + 1
					}
				} else {
					// haystackに有る文字で不一致となった場合はテキストの注目点を右にずらす
					// Note: tableの値は末尾で不一致となった場合のスキップ数のため、
					// 			見つかった文字数分スキップ数を減じる必要がある
					v := table[haystack[i+j]] - (needleLen - j - 1)
					if v <= 0 {
						// 後戻りしないように
						skip = 1
					} else {
						skip = v
					}
				}

				i += skip

				goto loop
			}
		}

		return i
	}

	return -1
}

func buidSkiptable(needle string) map[byte]int {
	table := make(map[byte]int)
	l := len(needle)

	for i := 0; i < l; i++ {
		table[needle[i]] = l - i - 1
	}

	return table
}

// Use builtin function
func strStr_Builtin(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// Brute force
// https://tech.pjin.jp/blog/2020/11/12/stringsearch-1/
func strStr_Bruteforce(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}

	for i := 0; i <= haystackLen-needleLen; i++ {
		if stringCompare(haystack[i:i+needleLen], needle) {
			return i
		}
	}

	return -1
}

func stringCompare(str1, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}
