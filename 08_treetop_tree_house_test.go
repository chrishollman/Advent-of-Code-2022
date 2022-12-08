package adventofcode2022

import "testing"

func TestDayEightChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Eight Challenge One example": {
			input: []string{"30373", "25512", "65332", "33549", "35390"},
			want:  21,
		},
		"Solve Day Eight Challenge One example two": {
			input: []string{"000", "999", "000"},
			want:  9,
		},
		"Solve Day Eight Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/08.txt")
				return res
			}(),
			want: 1684,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayEightChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayEightChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Eight Challenge Two example": {
			input: []string{"30373", "25512", "65332", "33549", "35390"},
			want:  8,
		},
		"Solve Day Eight Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/08.txt")
				return res
			}(),
			want: 486540,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayEightChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayEightChallengeOne(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/08.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayEightChallengeOne(input)
	}
}

func BenchmarkDayEightChallengeTwo(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/08.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayEightChallengeTwo(input)
	}
}
