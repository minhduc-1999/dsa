package longestconsecutivesequence

import "testing"

// case 0: [100,4,200,1,3,2]
// case 1: [0,3,7,2,5,8,4,6,0,1]
// case 2: [1,0,1,2]

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 0",
			args: args{
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			name: "case 1",
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			want: 9,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 0, 1, 2},
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{34, 5, 12, 9, 23, 7, 1, 40, 16, 3, 20, 8, 25, 2, 6, 18, 19, 37, 13, 30},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	arr := []int{100, 350, 200, 225, 227, 400, 56, 5, 3, 56, 789, 344, 12, 16, 25, 700, 505, 15, 300, 20, 11, 9, 10, 7, 2, 1000, 33, 12, 500, 6, 8, 100, 105, 110, 111, 1000, 130, 120, 2000, 112, 210, 600, 8, 23, 50, 47, 53, 34, 35, 1003, 450, 600, 30, 75, 3000, 72, 75, 201, 302, 901, 125, 5000, 999, 1001, 6000, 40, 32, 33, 3, 250, 503, 140, 1001, 4000, 300, 60, 700, 220, 9999, 160, 70, 20, 7000, 300, 510, 1005, 5, 55, 17, 88, 0, 150, 232, 455, 555, 420, 300, 24, 1123, 451, 202, 5555, 7777, 14, 2500, 153, 400, 12, 222, 111, 450, 899, 1020, 124, 231, 33, 1004, 304, 29, 1010, 152, 10000, 20, 133, 55, 10, 12, 25, 44, 1200, 332, 1500, 750, 234, 788, 12, 67, 17, 1201, 400, 11, 108, 99, 1040, 505, 21, 245, 700, 23, 231, 92, 600, 1300, 1, 10, 9, 11, 100, 200, 300, 50, 99, 250, 111, 333}
	for i := 0; i < b.N; i++ {
		longestConsecutive(arr)
	}
}

func BenchmarkLongestConsecutive2(b *testing.B) {
	arr := []int{100, 350, 200, 225, 227, 400, 56, 5, 3, 56, 789, 344, 12, 16, 25, 700, 505, 15, 300, 20, 11, 9, 10, 7, 2, 1000, 33, 12, 500, 6, 8, 100, 105, 110, 111, 1000, 130, 120, 2000, 112, 210, 600, 8, 23, 50, 47, 53, 34, 35, 1003, 450, 600, 30, 75, 3000, 72, 75, 201, 302, 901, 125, 5000, 999, 1001, 6000, 40, 32, 33, 3, 250, 503, 140, 1001, 4000, 300, 60, 700, 220, 9999, 160, 70, 20, 7000, 300, 510, 1005, 5, 55, 17, 88, 0, 150, 232, 455, 555, 420, 300, 24, 1123, 451, 202, 5555, 7777, 14, 2500, 153, 400, 12, 222, 111, 450, 899, 1020, 124, 231, 33, 1004, 304, 29, 1010, 152, 10000, 20, 133, 55, 10, 12, 25, 44, 1200, 332, 1500, 750, 234, 788, 12, 67, 17, 1201, 400, 11, 108, 99, 1040, 505, 21, 245, 700, 23, 231, 92, 600, 1300, 1, 10, 9, 11, 100, 200, 300, 50, 99, 250, 111, 333}
	for i := 0; i < b.N; i++ {
		longestConsecutive2(arr)
	}
}
