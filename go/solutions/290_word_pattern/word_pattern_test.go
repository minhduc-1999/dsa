package wordpattern

import "testing"

// case 0: "abba"	"dog cat cat dog"
// case 1: "abba"	"dog cat cat fish"
// case 2: "aaaa"	"dog cat cat dog"
func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				pattern: "abba",
				s:       "dog cat cat dog",
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				pattern: "abba",
				s:       "dog cat cat fish",
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				pattern: "aaaa",
				s:       "dog cat cat dog",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
