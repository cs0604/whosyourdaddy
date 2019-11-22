package decode_string

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				str: "3[ac4[a]]",
			},
			want: "acaaaaacaaaaacaaaa",
		},
		{
			name: "2",
			args: args{
				str: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "3",
			args: args{
				str: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "4",
			args: args{
				str: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
		{
			name: "5",
			args: args{
				str: "2[abc]ef3[cd]",
			},
			want: "abcabcefcdcdcd",
		},
		{
			name: "6",
			args: args{
				str: "3[a]2[b4[F]c]",
			},
			want: "aaabFFFFcbFFFFc",
		},
		{
			name: "7",
			args: args{
				str: "dd3[a]2[b4[F]c]ee",
			},
			want: "ddaaabFFFFcbFFFFcee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString_stack(tt.args.str); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
