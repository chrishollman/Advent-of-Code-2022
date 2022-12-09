package adventofcode2022

import "testing"

func TestDayNineChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Eight Challenge One example": {
			input: []string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"},
			want:  13,
		},
		"Solve Day Eight Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/09.txt")
				return res
			}(),
			want: -1,
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
