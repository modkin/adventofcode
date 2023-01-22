package main

import "testing"

func Test_wrapPos(t *testing.T) {
	type args struct {
		pos int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lowWrap",
			args: args{
				pos: 0,
				max: 5,
			},
			want: 5,
		},
		{
			name: "HighWrap",
			args: args{
				pos: 6,
				max: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapPos(tt.args.pos, tt.args.max); got != tt.want {
				t.Errorf("wrapPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
