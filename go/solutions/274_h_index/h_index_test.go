package hindex

import "testing"

// Testcases:
// case 0: [3,0,6,1,5]
// case 1: [1,3,1]
func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				citations: []int{3, 0, 6, 1, 5},
			},
			want: 3,
		},
		{
			name: "test 2",
			args: args{
				citations: []int{1, 3, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
