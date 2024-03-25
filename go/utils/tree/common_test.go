package tree

import (
	"reflect"
	"testing"
)

func TestNewFromArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Case 0",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			name: "Case 1",
			args: args{
				arr: []int{1, 2, 3, -9999, -9999, 4, 5},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "Case 2",
			args: args{
				arr: []int{},
			},
			want: nil,
		},
		{
			name: "Case 3",
			args: args{
				arr: []int{1, 2, 3, 6, -9999, 4, 5},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "Case 4",
			args: args{
				arr: []int{1, 2, 3, -9999, 6, 4, 5},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromArray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
