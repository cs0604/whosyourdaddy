package course_schedule

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{0, 1},
					{2, 3},
					{1, 2},
					{3, 0},
				},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{1, 0},
					{0, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
