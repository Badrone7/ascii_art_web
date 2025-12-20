package functions

import (
	"os"
	"strings"
	"unicode"
)

// the funtion that makes the art
func ArtMaker(text string, style string) ([]byte, error) {
	finalArt := []byte{}
	if text == "" {
		return nil, nil
	}
	art, err := ArtSelect(style)
	if err != nil {
		return nil, err
	}
	for _, char := range text {
		if !unicode.IsSpace(char) && !(char >= 32 && char <= 126) {
			finalArt = []byte("Error: Unsupported character detected.\n")
			return finalArt, nil
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
	return finalArt, nil
}

// a function that prints the art
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

// a function that selects the art type
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

// a function that generates the art from a txt file
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
