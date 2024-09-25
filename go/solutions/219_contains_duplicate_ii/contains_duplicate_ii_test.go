package containsduplicateii

import "testing"

// case 0: [1,2,3,1]	3
// case 1: [1,0,1,1]	1
// case 2: [1,2,3,1,2,3]	2
func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "test 1",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
