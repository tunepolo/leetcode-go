package longestcommonprefix

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, compareStr := range strs[1:] {
		prefixLength := len(prefix)
		if compStrLen := len(compareStr); compStrLen == 0 {
			return ""
		} else if compStrLen < prefixLength {
			prefixLength = compStrLen
			prefix = prefix[:prefixLength]
		}

		for i := 0; i < prefixLength; i++ {
			if prefix[i] != compareStr[i] {
				prefix = prefix[:i]
				break
			}
		}
	}

	return prefix
}
