package candy

import "testing"

// case 0: [1,0,2]
// case 1: [1,2,2]
// case 2: [1,2,1,0]
// case 4: [1,3,2,2,1]
// case 6: [1,3,4,5,2]
func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			name: "test 2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
		{
			name: "test 3",
			args: args{
				ratings: []int{1, 2, 1, 0},
			},
			want: 7,
		},
		{
			name: "test 4",
			args: args{
				ratings: []int{3, 2, 0},
			},
			want: 6,
		},
		{
			name: "Test 5",
			args: args{
				ratings: []int{1, 3, 2, 2, 1},
			},
			want: 7,
		},
		{
			name: "TEst 6",
			args: args{
				ratings: []int{1, 3, 4, 5, 2},
			},
			want: 11,
		},
		{
			name: "Test 7",
			args: args{
				ratings: []int{1, 2, 3, 3, 2, 1},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
