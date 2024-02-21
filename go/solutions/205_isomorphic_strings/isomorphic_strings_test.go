package isomorphicstrings

import "testing"

// Testcases:
// case 0: "egg"	"add"
// case 1: "foo"	"bar"
// case 2: "paper"	"title"
func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
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
        s: "egg",
        t: "add",
      },
    },
    {
      name: "case 1",
      want: false,
      args: args{
        s: "foo",
        t: "bar",
      },
    },
    {
      name: "case 2",
      want: true,
      args: args{
        s: "paper",
        t: "title",
      },
    },
    {
      name: "case 3",
      want: false,
      args: args{
        s: "badc",
        t: "baba",
      },
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
