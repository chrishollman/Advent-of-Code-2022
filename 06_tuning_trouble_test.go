package adventofcode2022

import (
	"testing"
)

func TestDaySixChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringWantInt{
		"Solve Day Six Challenge One example": {
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  7,
		},
		"Solve Day Six Challenge One actual": {
			input: func() string {
				res, _ := readFileIntoStringArray("./input/06.txt")
				return res[0]
			}(),
			want: 1929,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := daySixChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDaySixChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringWantInt{
		"Solve Day Six Challenge Two example": {
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  19,
		},
		"Solve Day Six Challenge Two actual": {
			input: func() string {
				res, _ := readFileIntoStringArray("./input/06.txt")
				return res[0]
			}(),
			want: 3298,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := daySixChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDaySixChallengeOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		daySixChallengeOne("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	}
}

func BenchmarkDaySixChallengeTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		daySixChallengeTwo("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	}
}
