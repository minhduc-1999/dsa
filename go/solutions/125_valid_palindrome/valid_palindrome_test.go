package validpalindrome

import "testing"

// Testcases:
// case 0: "A man, a plan, a canal: Panama"
// case 1: "race a car"
// case 2: " "
func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
    {
      name: "case 0",
      args: args{
        s: "A man, a plan, a canal: Panama",
      },
      want: true,
    },
    {
      name: "case 1",
      args: args{
        s: "race a car",
      },
      want: false,
    },
    {
      name: "case 2",
      args: args{
        s: " ",
      },
      want: true,
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
