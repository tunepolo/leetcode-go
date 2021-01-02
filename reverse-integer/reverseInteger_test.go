package reverseinteger

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "Negative integer",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "0 end",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "Bigger than threshold",
			args: args{x: 1534236469},
			want: 0,
		},
		{
			name: "Smaller than threshold",
			args: args{x: -1147483647},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
