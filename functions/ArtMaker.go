package functions

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// the function that will check if the text contains only new lines
func Onlynewlines(text string) int {
	count := 0
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if runes[i] != '\\' {
			return -1
		}
		if i+1 >= len(runes) || runes[i+1] != 'n' {
			return -1
		}
		count++
		if i+2 >= len(runes) {
			break
		}
		runes = runes[i+2:]
		i = -1
	}
	return count
}

// the funtion that makes the art
func ArtMaker(text string, style string) []byte {
	finalArt := []byte{}
	if text == "" {
		return nil
	}
	art := ArtSelect(style)
	for _, char := range text {
		if !unicode.IsSpace(char) && !(char >= 32 && char <= 126) {
			finalArt = []byte("Error: Unsupported character detected.\n")
			return finalArt
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
	return finalArt
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
func ArtSelect(style string) [][]string {
	style = "resources/" + style + ".txt"
	file, err := os.ReadFile(style)
	if err != nil {
		fmt.Println("Error reading file:", err)
		
	}
	if style == "resources/thinkertoy.txt" {
		file = []byte(strings.ReplaceAll(string(file), "\r\n", "\n"))
	}
	art := ArtGenerator(file)
	return art
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
