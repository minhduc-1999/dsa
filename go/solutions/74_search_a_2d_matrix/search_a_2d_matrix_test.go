package searcha2dmatrix

import "testing"

// Testcases:
// case 0: [[1,3,5,7],[10,11,16,20],[23,30,34,60]]	3
// case 1: [[1,3,5,7],[10,11,16,20],[23,30,34,60]]	13
func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 3,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 13,
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{{1}},
				target: 2,
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				matrix: [][]int{{1, 1}},
				target: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
