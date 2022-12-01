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
			want: 24000,
			input: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
		},
		{
			name:  "actual",
			want:  67027,
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
			want: 45000,
			input: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
		},
		{
			name:  "actual",
			want:  197291,
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
