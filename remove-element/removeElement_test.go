package removeElement

import (
	"reflect"
	"sort"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: args{
				nums: []int{2, 2, 0, 0},
				val:  2,
			},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: args{
				nums: []int{0, 1, 4, 0, 3, 0, 0, 0},
				val:  5,
			},
		},
		{
			name: "Empty input",
			args: args{
				nums: []int{},
				val:  0,
			},
			want: args{
				nums: []int{},
				val:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			in := append([]int{}, tt.args.nums[0:got]...)
			out := append([]int{}, tt.want.nums[0:got]...)
			sort.Ints(in)
			sort.Ints(out)
			if got != tt.want.val || !reflect.DeepEqual(in, out) {
				t.Errorf("removeElement() = (%v, %v), want %v", in, got, out)
			}
		})
	}
}
