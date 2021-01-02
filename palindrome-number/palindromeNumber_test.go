package palindromenumber

import (
	"testing"
)

type args struct {
	x int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "Case 1",
		args: args{x: 121},
		want: true,
	},
	{
		name: "Case 2",
		args: args{x: -121},
		want: false,
	},
	{
		name: "Case 3",
		args: args{x: 10},
		want: false,
	},
	{
		name: "Case 4",
		args: args{x: 1221},
		want: true,
	},
	{
		name: "Case 5",
		args: args{x: 12321},
		want: true,
	},
}

func Test_isPalindrome(t *testing.T) {
	tests := tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome2(t *testing.T) {
	tests := tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
