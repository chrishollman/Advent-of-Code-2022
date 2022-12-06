package adventofcode2022

import (
	"testing"
)

func TestDayFiveChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantString{
		"Solve Day Five Challenge One example": {
			input: []string{"    [D]    ", "[N] [C]    ", "[Z] [M] [P]", " 1   2   3 ", "", "move 1 from 2 to 1", "move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"},
			want:  "CMZ",
		},
		"Solve Day Five Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/05.txt")
				return res
			}(),
			want: "DHBJQJCCW",
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayFiveChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayFiveChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantString{
		"Solve Day Five Challenge One example": {
			input: []string{"    [D]    ", "[N] [C]    ", "[Z] [M] [P]", " 1   2   3 ", "", "move 1 from 2 to 1", "move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"},
			want:  "MCD",
		},
		"Solve Day Five Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/05.txt")
				return res
			}(),
			want: "WJVRLSJJT",
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayFiveChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayFiveChallengeOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayFiveChallengeOne([]string{"    [D]    ", "[N] [C]    ", "[Z] [M] [P]", " 1   2   3 ", "", "move 1 from 2 to 1", "move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"})
	}
}

func BenchmarkDayFiveChallengeTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayFiveChallengeTwo([]string{"    [D]    ", "[N] [C]    ", "[Z] [M] [P]", " 1   2   3 ", "", "move 1 from 2 to 1", "move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"})
	}
}
