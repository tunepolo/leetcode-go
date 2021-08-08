package impl

import "testing"

type args struct {
	haystack string
	needle   string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "Example1",
		args: args{
			haystack: "hello",
			needle:   "ll",
		},
		want: 2,
	},
	{
		name: "Example2",
		args: args{
			haystack: "aaaaa",
			needle:   "bba",
		},
		want: -1,
	},
	{
		name: "Example3",
		args: args{
			haystack: "",
			needle:   "",
		},
		want: 0,
	},
	{
		name: "empty haystack",
		args: args{
			haystack: "",
			needle:   "a",
		},
		want: -1,
	},
	{
		name: "haystack and needle are equal",
		args: args{
			haystack: "a",
			needle:   "a",
		},
		want: 0,
	},
	{
		name: "Case failed in leetcode submit check 1",
		args: args{
			haystack: "mississippi",
			needle:   "issi",
		},
		want: 1,
	},
	{
		name: "Case failed in leetcode submit check 2",
		args: args{
			haystack: "ababbbbaaabbbaaa",
			needle:   "bbbb",
		},
		want: 3,
	},
	{
		name: "Case failed in leetcode submit check 3",
		args: args{
			haystack: "mississippi",
			needle:   "issipi",
		},
		want: -1,
	},
	{
		name: "Case failed in leetcode submit check 4",
		args: args{
			haystack: "bbbbababbbaabbba",
			needle:   "abb",
		},
		want: 6,
	},
	{
		name: "Case failed in leetcode submit check 5",
		args: args{
			haystack: "baabbaaaaaaabbaaaaabbabbababaabbabbbbbabbabbbbbbabababaabbbbbaaabbbbabaababababbbaabbbbaaabbaababbbaabaabbabbaaaabababaaabbabbababbabbaaabbbbabbbbabbabbaabbbaa",
			needle:   "bbaaaababa",
		},
		want: 107,
	},
	{
		name: "Case failed in leetcode submit check 6",
		args: args{
			haystack: "aabaabbbaabbbbabaaab",
			needle:   "abaa",
		},
		want: 1,
	},
}

func Test_strStr(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr_BM(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr_BM(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr_Builtin(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr_Builtin(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr_Bruteforce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr_Bruteforce(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
