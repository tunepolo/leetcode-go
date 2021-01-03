package validparentheses

// https://leetcode.com/problems/valid-parentheses/

import (
	"strings"
)

func isValid(s string) bool {
	var stack []string
	for _, c := range strings.Split(s, "") {
		sl := len(stack)
		if sl != 0 {
			switch last := stack[sl-1]; {
			case last == "(" && c == ")":
				fallthrough
			case last == "{" && c == "}":
				fallthrough
			case last == "[" && c == "]":
				stack = stack[:sl-1] // Stack pop
				continue
			}
		}
		stack = append(stack, c) // Stack push
	}

	return len(stack) == 0
}
