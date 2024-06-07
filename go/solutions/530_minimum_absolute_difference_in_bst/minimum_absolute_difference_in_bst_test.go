package minimumabsolutedifferenceinbst

import (
	. "dsa/utils/tree"
	"testing"
)

// Testcases:
// case 0: [4,2,6,1,3]
// case 1: [1,0,48,null,null,12,49]
// [236,104,701,null,227,null,911]
func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: NewFromArray([]int{4, 2, 6, 1, 3}),
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				root: NewFromArray([]int{1, 0, 48, -9999, -9999, 12, 49}),
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				root: NewFromArray([]int{236, 104, 701, -9999, 227, -9999, 911}),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
