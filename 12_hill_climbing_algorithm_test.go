package adventofcode2022

import (
	"testing"
)

func TestDayTwelveChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Twelve Challenge One example": {
			input: []string{"Sabqponm", "abcryxxl", "accszExk", "acctuvwj", "abdefghi"},
			want:  31,
		},
		"Solve Day Twelve Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/12.txt")
				return res
			}(),
			want: 330,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayTwelveChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayTwelveChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Twelve Challenge One example": {
			input: []string{"Sabqponm", "abcryxxl", "accszExk", "acctuvwj", "abdefghi"},
			want:  29,
		},
		"Solve Day Twelve Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/12.txt")
				return res
			}(),
			want: 321,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayTwelveChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayTwelveChallengeOne(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/12.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayTwelveChallengeOne(input)
	}
}

func BenchmarkDayTwelveChallengeTwo(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/12.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayTwelveChallengeTwo(input)
	}
}
