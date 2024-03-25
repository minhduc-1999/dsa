package binarytreelevelordertraversal

import (
	. "dsa/utils/tree"
	"reflect"
	"testing"
)

// Testcases:
// case 0: [3,9,20,null,null,15,7]
// case 1: [1]
// case 2: []
func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Case 0",
			args: args{
				root: NewFromArray([]int{3, 9, 20, -9999, -9999, 15, 7}),
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			name: "Case 1",
			args: args{
				root: NewFromArray([]int{1}),
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "Case 2",
			args: args{
				root: NewFromArray([]int{}),
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
