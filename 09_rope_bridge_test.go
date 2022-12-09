package adventofcode2022

import (
	"testing"
)

func TestDayNineChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Nine Challenge One example": {
			input: []string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"},
			want:  13,
		},
		"Solve Day Nine Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/09.txt")
				return res
			}(),
			want: 6486,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayNineChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayNineChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Nine Challenge Two example": {
			input: []string{"R 5", "U 8", "L 8", "D 3", "R 17", "D 10", "L 25", "U 20"},
			want:  36,
		},
		"Solve Day Nine Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/09.txt")
				return res
			}(),
			want: 2678,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayNineChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayNineChallengeOne(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/09.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayNineChallengeOne(input)
	}
}

func BenchmarkDayNineChallengeTwo(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/09.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayNineChallengeTwo(input)
	}
}
