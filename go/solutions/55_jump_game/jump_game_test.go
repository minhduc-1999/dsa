package jumpgame

import "testing"

// Testcases:
// case 0: [2,3,1,1,4]
// case 1: [3,2,1,0,4]
func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 0",
			want: true,
			args: args{
				nums: []int{
					2, 3, 1, 1, 4,
				},
			},
		},
		{
			name: "case 1",
			want: false,
			args: args{
				nums: []int{
					3, 2, 1, 0, 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
