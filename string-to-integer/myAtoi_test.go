package myatoi

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "Example2",
			args: args{s: "   -42"},
			want: -42,
		},
		{
			name: "Example3",
			args: args{s: "4193 with words"},
			want: 4193,
		},
		{
			name: "Example4",
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			name: "Example5",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "int64 upper bound",
			args: args{s: "9223372036854775808"},
			want: 2147483647,
		},
		{
			name: "Digits after words",
			args: args{s: "123 abc 456"},
			want: 123,
		},
		{
			name: "Plus Minus both used",
			args: args{s: "+-12"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterInteger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Only digits",
			args: args{s: "42"},
			want: "42",
		},
		{
			name: "With space and minus",
			args: args{s: "   -42"},
			want: "-42",
		},
		{
			name: "Digits before words",
			args: args{s: "4193 with words"},
			want: "4193",
		},
		{
			name: "Digits after words",
			args: args{s: "words and 987"},
			want: "",
		},
		{
			name: "Digits before / after words",
			args: args{s: "123 abc 456"},
			want: "123",
		},
		{
			name: "Plus Minus both used",
			args: args{s: "+-12"},
			want: "",
		},
		{
			name: "Plus with words",
			args: args{s: "+sample"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterInteger(tt.args.s); got != tt.want {
				t.Errorf("findInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
