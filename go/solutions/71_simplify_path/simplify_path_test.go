package simplifypath

import "testing"

// Testcases:
// case 0: "/home/"
// case 1: "/../"
// case 2: "/home//foo/"
func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
    {
      name: "case 0",
      want: "/home",
      args: args{
        path: "/home/",
      },
    },
    {
      name: "case 1",
      want: "/",
      args: args{
        path: "/../",
      },
    },
    {
      name: "case 2",
      want: "/home/foo",
      args: args{
        path: "/home//foo/",
      },
    },
    {
      name: "case 3",
      want: "/a/b/c",
      args: args{
        path: "/a//b////c/d//././/..",
      },
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
