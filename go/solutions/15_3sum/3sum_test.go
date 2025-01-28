package threesum

import (
	"reflect"
	"testing"
)

// case 0: [-1,0,1,2,-1,-4]
// case 1: [0,1,1]
// case 2: [0,0,0]
func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{-2, 0, 0, 2, 2},
			},
			want: [][]int{{-2, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
