package word_search

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCE",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "SEE",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCD",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				board: [][]byte{
					{'A', 'A', 'C', 'E'},
					{'A', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "AAAA",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
