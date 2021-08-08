package searchInsertPosition

import "testing"

type args struct {
	nums   []int
	target int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "Example1",
		args: args{
			nums:   []int{1, 3, 5, 6},
			target: 5,
		},
		want: 2,
	},
	{
		name: "Example2",
		args: args{
			nums:   []int{1, 3, 5, 6},
			target: 2,
		},
		want: 1,
	},
	{
		name: "Example3",
		args: args{
			nums:   []int{1, 3, 5, 6},
			target: 7,
		},
		want: 4,
	},
	{
		name: "Example4",
		args: args{
			nums:   []int{1, 3, 5, 6},
			target: 0,
		},
		want: 0,
	},
	{
		name: "Example5",
		args: args{
			nums:   []int{1},
			target: 0,
		},
		want: 0,
	},
}

func Test_searchInsert(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInsert_Bruteforce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert_Bruteforce(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert_Bruteforce() = %v, want %v", got, tt.want)
			}
		})
	}
}
