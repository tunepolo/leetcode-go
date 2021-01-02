package romantointeger

import (
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple case 1",
			args: args{s: "II"},
			want: 2,
		},
		{
			name: "Simple case 2",
			args: args{s: "XII"},
			want: 12,
		},
		{
			name: "Simple case 3",
			args: args{s: "XXVII"},
			want: 27,
		},
		{
			name: "Need subtraction 1",
			args: args{s: "IV"},
			want: 4,
		},
		{
			name: "Need subtraction 2",
			args: args{s: "IX"},
			want: 9,
		},
		{
			name: "Need subtraction 3",
			args: args{s: "XL"},
			want: 40,
		},
		{
			name: "Need subtraction 4",
			args: args{s: "XC"},
			want: 90,
		},
		{
			name: "Need subtraction 5",
			args: args{s: "CD"},
			want: 400,
		},
		{
			name: "Need subtraction 6",
			args: args{s: "CM"},
			want: 900,
		},
		{
			name: "Complex case 1",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "Complex case 2",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
