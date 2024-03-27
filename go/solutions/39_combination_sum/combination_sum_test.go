package combinationsum

import (
	"reflect"
	"testing"
)

// Testcases:
// case 0: [2,3,6,7]	7
// case 1: [2,3,5]	8
// case 2: [2]	1
func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	name: "Case 0",
		// 	args: args{
		// 		candidates: []int{2, 3, 6, 7},
		// 		target:     7,
		// 	},
		// 	want: [][]int{{2, 2, 3}, {7}},
		// },
		// {
		// 	name: "Case 1",
		// 	args: args{
		// 		candidates: []int{2, 3, 5},
		// 		target:     8,
		// 	},
		// 	want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		// },
		// {
		// 	name: "Case 2",
		// 	args: args{
		// 		candidates: []int{2},
		// 		target:     1,
		// 	},
		// 	want: [][]int{},
		// },
		{
			name: "Case 3",
			args: args{
				candidates: []int{8, 7, 4, 3},
				target:     11,
			},
			want: [][]int{{8, 3}, {7, 4}, {4, 4, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
