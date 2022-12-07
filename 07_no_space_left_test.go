package adventofcode2022

import "testing"

func TestDaySevenChallengeOne(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Seven Challenge One example": {
			input: []string{"$ cd /", "$ ls", "dir a", "14848514 b.txt", "8504156 c.dat", "dir d", "$ cd a", "$ ls", "dir e", "29116 f", "2557 g", "62596 h.lst", "$ cd e", "$ ls", "584 i", "$ cd ..", "$ cd ..", "$ cd d", "$ ls", "4060174 j", "8033020 d.log", "5626152 d.ext", "7214296 k"},
			want:  95437,
		},
		"Solve Day Seven Challenge One actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/07.txt")
				return res
			}(),
			want: 1297683,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := daySevenChallengeOne(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestDaySevenChallengeTwo(t *testing.T) {
	tests := map[string]TestConfigStringSliceWantInt{
		"Solve Day Seven Challenge Two example": {
			input: []string{"$ cd /", "$ ls", "dir a", "14848514 b.txt", "8504156 c.dat", "dir d", "$ cd a", "$ ls", "dir e", "29116 f", "2557 g", "62596 h.lst", "$ cd e", "$ ls", "584 i", "$ cd ..", "$ cd ..", "$ cd d", "$ ls", "4060174 j", "8033020 d.log", "5626152 d.ext", "7214296 k"},
			want:  24933642,
		},
		"Solve Day Seven Challenge Two actual": {
			input: func() []string {
				res, _ := readFileIntoStringArray("./input/07.txt")
				return res
			}(),
			want: 5756764,
		},
	}

	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			got := daySevenChallengeTwo(cfg.input)
			want := cfg.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
