package permutations

import (
	"reflect"
	"testing"
)

// Testcases:
// case 0: [1,2,3]
// case 1: [0,1]
// case 2: [1]

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1},
			},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
