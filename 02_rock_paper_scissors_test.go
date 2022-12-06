package adventofcode2022

import "testing"

func TestDayTwoChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Two Challenge One example": {
			input: []string{"A Y", "B X", "C Z"},
			want:  15,
		},
		"Solve Day Two Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/02.txt")
				return res
			}(),
			want: 11666,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayTwoChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayTwoChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Two Challenge Two example": {
			input: []string{"A Y", "B X", "C Z"},
			want:  12,
		},
		"Solve Day Two Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/02.txt")
				return res
			}(),
			want: 12767,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayTwoChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayTwoChallengeOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayTwoChallengeOne([]string{"A Y", "B X", "C Z"})
	}
}

func BenchmarkDayTwoChallengeTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayTwoChallengeTwo([]string{"A Y", "B X", "C Z"})
	}
}
