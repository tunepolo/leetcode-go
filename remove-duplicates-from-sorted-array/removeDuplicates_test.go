package removeduplicates

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type want struct {
		ret  int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: want{
				ret:  2,
				nums: []int{1, 2},
			},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: want{
				ret:  5,
				nums: []int{0, 1, 2, 3, 4},
			},
		},
		{
			name: "Empty Input",
			args: args{
				nums: []int{},
			},
			want: want{
				ret:  0,
				nums: []int{},
			},
		},
		{
			name: "1 Item Input",
			args: args{
				nums: []int{1},
			},
			want: want{
				ret:  1,
				nums: []int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want.ret || !reflect.DeepEqual(tt.args.nums[0:got], tt.want.nums) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
