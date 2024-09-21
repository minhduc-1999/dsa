package addbinary

import "testing"

// Testcases:
// case 0: "11"	"1"
// case 1: "1010"	"1011"
func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		// {
		// 	name: "case 2",
		// 	args: args{
		// 		a: "1010",
		// 		b: "1011",
		// 	},
		// 	want: "10101",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
