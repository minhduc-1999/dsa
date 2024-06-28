package convertsortedarraytobinarysearchtree

import (
	. "dsa/utils/tree"
	"reflect"
	"testing"
)

// case 0: [-10,-3,0,5,9]
// case 1: [1,3]
func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: NewFromArray([]int{0, -10, 5, -9999, -3, -9999, 9}),
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 3},
			},
			want: NewFromArray([]int{1, -9999, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot \t%v\nwant\t%v", PrintTreeBFS(got), PrintTreeBFS(tt.want))
			}
		})
	}
}
