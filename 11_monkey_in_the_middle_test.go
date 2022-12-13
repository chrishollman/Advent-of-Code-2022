package adventofcode2022

import (
	"testing"
)

func TestDayElevenChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Eleven Challenge One example": {
			input: []string{"Monkey 0:", "  Starting items: 79, 98", "  Operation: new = old * 19", "  Test: divisible by 23", "    If true: throw to monkey 2", "    If false: throw to monkey 3", " ", "Monkey 1:", "  Starting items: 54, 65, 75, 74", "  Operation: new = old + 6", "  Test: divisible by 19", "    If true: throw to monkey 2", "    If false: throw to monkey 0", " ", "Monkey 2:", "  Starting items: 79, 60, 97", "  Operation: new = old * old", "  Test: divisible by 13", "    If true: throw to monkey 1", "    If false: throw to monkey 3", " ", "Monkey 3:", "  Starting items: 74", "  Operation: new = old + 3", "  Test: divisible by 17", "    If true: throw to monkey 0", "    If false: throw to monkey 1"},
			want:  10605,
		},
		"Solve Day Eleven Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/11.txt")
				return res
			}(),
			want: 117640,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayElevenChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDayElevenChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Eleven Challenge Two example": {
			input: []string{"Monkey 0:", "  Starting items: 79, 98", "  Operation: new = old * 19", "  Test: divisible by 23", "    If true: throw to monkey 2", "    If false: throw to monkey 3", " ", "Monkey 1:", "  Starting items: 54, 65, 75, 74", "  Operation: new = old + 6", "  Test: divisible by 19", "    If true: throw to monkey 2", "    If false: throw to monkey 0", " ", "Monkey 2:", "  Starting items: 79, 60, 97", "  Operation: new = old * old", "  Test: divisible by 13", "    If true: throw to monkey 1", "    If false: throw to monkey 3", " ", "Monkey 3:", "  Starting items: 74", "  Operation: new = old + 3", "  Test: divisible by 17", "    If true: throw to monkey 0", "    If false: throw to monkey 1"},
			want:  2713310158,
		},
		"Solve Day Eleven Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/11.txt")
				return res
			}(),
			want: 30616425600,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := dayElevenChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDayElevenChallengeOne(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/11.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayElevenChallengeOne(input)
	}
}

func BenchmarkDayElevenChallengeTwo(b *testing.B) {
	input := func() []string {
		res, _ := readFileIntoStringArray("./input/11.txt")
		return res
	}()
	for n := 0; n < b.N; n++ {
		dayElevenChallengeTwo(input)
	}
}
