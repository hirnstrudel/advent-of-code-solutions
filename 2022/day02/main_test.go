package main

import "testing"

func Test_part1(t *testing.T) {

	tests := []struct {
		name  string
		want  int
		input string
	}{
		{
			name: "example",
			want: 15,
			input: `A Y
B X
C Z`,
		},
		{
			name:  "actual",
			want:  13484,
			input: input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("%v: part1() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {

	tests := []struct {
		name  string
		want  int
		input string
	}{
		{
			name: "example",
			want: 12,
			input: `A Y
B X
C Z`,
		},
		{
			name:  "actual",
			want:  13433,
			input: input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("%v: part2() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
