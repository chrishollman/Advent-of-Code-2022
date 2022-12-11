package adventofcode2022

import (
	"testing"
)

func TestDayTenChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Ten Challenge One example": {
			input: []string{"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8", "addx 13", "addx 4", "noop", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx -35", "addx 1", "addx 24", "addx -19", "addx 1", "addx 16", "addx -11", "noop", "noop", "addx 21", "addx -15", "noop", "noop", "addx -3", "addx 9", "addx 1", "addx -3", "addx 8", "addx 1", "addx 5", "noop", "noop", "noop", "noop", "noop", "addx -36", "noop", "addx 1", "addx 7", "noop", "noop", "noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop", "noop", "addx 1", "noop", "noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13", "addx 7", "noop", "addx 1", "addx -33", "noop", "noop", "noop", "addx 2", "noop", "noop", "noop", "addx 8", "noop", "addx -1", "addx 2", "addx 1", "noop", "addx 17", "addx -9", "addx 1", "addx 1", "addx -3", "addx 11", "noop", "noop", "addx 1", "noop", "addx 1", "noop", "noop", "addx -13", "addx -19", "addx 1", "addx 3", "addx 26", "addx -30", "addx 12", "addx -1", "addx 3", "addx 1", "noop", "noop", "noop", "addx -9", "addx 18", "addx 1", "addx 2", "noop", "noop", "addx 9", "noop", "noop", "noop", "addx -1", "addx 2", "addx -37", "addx 1", "addx 3", "noop", "addx 15", "addx -21", "addx 22", "addx -6", "addx 1", "noop", "addx 2", "addx 1", "noop", "addx -10", "noop", "noop", "addx 20", "addx 1", "addx 2", "addx 2", "addx -6", "addx -11", "noop", "noop", "noop"},
			want:  13140,
		},
		"Solve Day Ten Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/10.txt")
				return res
			}(),
			want: 14240,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayTenChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayTenChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Ten Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/10.txt")
				return res
			}(),
			want: -1,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayTenChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayTenChallengeOne(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/10.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayTenChallengeOne(input)
	}
}

func BenchmarkDayTenChallengeTwo(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/10.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayTenChallengeTwo(input)
	}
}
