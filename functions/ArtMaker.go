package functions

import (
	"os"
	"strings"
	"unicode"
)

// ArtMaker generates ASCII art based on the provided text and style
func ArtMaker(text string, style string) ([]byte, error, int) {
	finalArt := []byte{}
	if text == "" {
		return nil, nil, 0
	}
	art, err := ArtSelect(style)
	if err != nil {
		return nil, err, 0
	}
	for _, char := range text {
		if !unicode.IsSpace(char) && !(char >= 32 && char <= 126) {
			finalArt = []byte("Error: Unsupported character detected.\n")
			return finalArt, nil, 1
		}
	}
	text = strings.ReplaceAll(text, "\r\n", "\n")
	txt := strings.Split(string(text), "\n")
	for i := 0; i < len(txt); i++ {
		if txt[i] == "" {
			finalArt = append(finalArt, '\n')
		} else {
			finalArt = append(finalArt, PrintArt(txt[i], art)...)
		}
	}
	return finalArt, nil, 0
}

// PrintArt constructs the ASCII art for a given line of text
func PrintArt(text string, art [][]string) []byte {
	finalArt := []byte{}
	for line := 0; line < 8; line++ {
		for _, char := range text {
			index := int(char) - 32
			finalArt = append(finalArt, []byte(art[index][line])...)
		}
		finalArt = append(finalArt, '\n')
	}
	return finalArt
}

// ArtSelect selects and reads the appropriate ASCII art style file
func ArtSelect(style string) ([][]string, error) {
	style = "resources/" + style + ".txt"
	file, err := os.ReadFile(style)
	if err != nil {
		return nil, err
	}
	if style == "resources/thinkertoy.txt" {
		file = []byte(strings.ReplaceAll(string(file), "\r\n", "\n"))
	}
	art := ArtGenerator(file)
	return art, nil
}

// ArtGenerator processes the ASCII art file into a structured format
func ArtGenerator(file []byte) [][]string {
	if file[0] == '\n' {
		file = file[1:]
	}
	count := 0
	character := []string{}
	lines := []rune{}
	art := [][]string{}
	for _, c := range string(file) {
		if count == 8 {
			art = append(art, character)
			count = 0
			character = []string{}
			continue
		}
		if c == '\n' {
			count++
			character = append(character, string(lines))
			lines = []rune{}
			continue
		}
		lines = append(lines, c)
	}
	if len(character) > 0 {
		art = append(art, character)
	}
	return art
}
