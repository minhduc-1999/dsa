package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
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
      want: true,
      args: args{
        s: "()",
      },
    },
    {
      name: "case 1",
      want: true,
      args: args{
        s: "()[]{}",
      },
    },
    {
      name: "case 2",
      want: false,
      args: args{
        s: "(]",
      },
    },
    {
      name: "case 3",
      want: true,
      args: args{
        s: "([{}]())",
      },
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
