package binarytreezigzaglevelordertraversal

import (
	. "dsa/utils/tree"
	"reflect"
	"testing"
)

// Testcases:
// case 0: [3,9,20,null,null,15,7]
// case 1: [1]
// case 2: []
func Test_zigzagLevelOrder(t *testing.T) {
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
				{20, 9},
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
		{
			name: "Case 3",
			args: args{
				root: NewFromArray([]int{1, 2, 3, 4, -9999, -9999, 5}),
			},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
