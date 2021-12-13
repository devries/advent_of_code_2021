package utils

import (
	"strings"
)

var alphabet = map[string]rune{
	".##.\n#..#\n#..#\n####\n#..#\n#..#": 'A',
	"###.\n#..#\n###.\n#..#\n#..#\n###.": 'B',
	".##.\n#..#\n#...\n#...\n#..#\n.##.": 'C',
	"####\n#...\n###.\n#...\n#...\n####": 'E',
	"####\n#...\n###.\n#...\n#...\n#...": 'F',
	".##.\n#..#\n#...\n#.##\n#..#\n.###": 'G',
	"#..#\n#..#\n####\n#..#\n#..#\n#..#": 'H',
	".###\n..#.\n..#.\n..#.\n..#.\n.###": 'I',
	"..##\n...#\n...#\n...#\n#..#\n.##.": 'J',
	"#..#\n#.#.\n##..\n#.#.\n#.#.\n#..#": 'K',
	"#...\n#...\n#...\n#...\n#...\n####": 'L',
	".##.\n#..#\n#..#\n#..#\n#..#\n.##.": 'O',
	"###.\n#..#\n#..#\n###.\n#...\n#...": 'P',
	"###.\n#..#\n#..#\n###.\n#.#.\n#..#": 'R',
	".###\n#...\n#...\n.##.\n...#\n###.": 'S',
	"#..#\n#..#\n#..#\n#..#\n#..#\n.##.": 'U',
	"#...\n#...\n.#.#\n..#.\n..#.\n..#.": 'Y',
	"####\n...#\n..#.\n.#..\n#...\n####": 'Z',
}

// Convert Advent of Code height 6 font letters to a string.
// The message must be represented as 6 rows of # and . characters
// separated by newlines.
func OCRLetters(message string) string {
	lines := strings.Split(strings.TrimSpace(message), "\n")
	letters := []rune{}

	for i := 0; i < len(lines[0]); i += 5 {
		letterLines := []string{}
		for _, ln := range lines {
			letterLines = append(letterLines, ln[i:i+4])
		}
		letter := strings.Join(letterLines, "\n")
		letters = append(letters, alphabet[letter])
	}

	return string(letters)
}
