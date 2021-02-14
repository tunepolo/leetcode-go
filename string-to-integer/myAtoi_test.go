package myatoi

import "testing"

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
			name: "Digits after words",
			args: args{s: "123 abc 456"},
			want: 123,
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
