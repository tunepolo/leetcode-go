package longestsubstringwithoutrepeatingcharacters

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	answer := 0
	charMap := make(map[byte]int)

	// i:走査の先頭位置, j:走査中の位置
	for i, j := 0, 0; j < len(s); j++ {
		charMap[s[j]]++

		// 重複した場合、その文字が無くなるまで先頭位置をずらす
		for charMap[s[j]] == 2 && i < j {
			charMap[s[i]]--
			i++
		}

		if answer < (j - i + 1) {
			answer = (j - i + 1)
		}
	}

	return answer
}
