package adventofcode2022

import "testing"

func TestDayFourChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Four Challenge One example": {
			input: []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"},
			want:  2,
		},
		"Solve Day Four Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/04.txt")
				return res
			}(),
			want: 500,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayFourChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayFourChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Four Challenge Two example": {
			input: []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"},
			want:  4,
		},
		"Solve Day Four Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/04.txt")
				return res
			}(),
			want: 815,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayFourChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayFourChallengeOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayFourChallengeOne([]string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"})
	}
}

func BenchmarkDayFourChallengeTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayFourChallengeTwo([]string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"})
	}
}
