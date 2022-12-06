package adventofcode2022

import (
	"testing"
)

func TestDayOneChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day One Challenge One example": {
			input: []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000", ""},
			want:  24000,
		},
		"Solve Day One Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/01.txt")
				return res
			}(),
			want: 74198,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayOneChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayOneChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day One Challenge Two example": {
			input: []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000", ""},
			want:  45000,
		},
		"Solve Day One Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/01.txt")
				return res
			}(),
			want: 209914,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayOneChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayOneChallengeOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayOneChallengeOne([]string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000", ""})
	}
}

func BenchmarkDayOneChallengeTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayOneChallengeTwo([]string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000", ""})
	}
}
