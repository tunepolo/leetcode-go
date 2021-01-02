package longestcommonprefix

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Case 2",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "Case 3",
			args: args{
				strs: []string{},
			},
			want: "",
		},
		{
			name: "Case 4",
			args: args{
				strs: []string{"flower"},
			},
			want: "flower",
		},
		{
			name: "Case 5",
			args: args{
				strs: []string{"flower", "flow", ""},
			},
			want: "",
		},
		{
			name: "Case 6",
			args: args{
				strs: []string{"ab", "a"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
