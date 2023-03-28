package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PrintAscii returns a banner of the given text in ASCII art format.
// The style parameter determines the style of the banner, which can be one of:
// "standard", "thinkertoy", or "shadow".
func PrintAscii(text, style string) (string, error) {
	// Determine the path of the banner file based on the style parameter.
	path := ""
	switch style {
	case "standard":
		path = "banners/standard.txt"
	case "thinkertoy":
		path = "banners/thinkertoy.txt"
	case "shadow":
		path = "banners/shadow.txt"
	default:
		return "", fmt.Errorf("unsupported banner style: %s", style)
	}

	// Read the banner file.
	data, err := ReadBanner(path)
	if err != nil {
		return "", fmt.Errorf("cannot read banner file: %w", err)
	}
	if len(data) == 0 {
		return "", fmt.Errorf("banner file is empty: %s", path)
	}

	// Convert the text into a banner by mapping each character to a block of ASCII art.
	var out strings.Builder
	words := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	for i, word := range words {
		if i > 0 && word == "" {
			out.WriteString("\n")
		}
		if word == "" {
			continue
		}
		for j := 0; j < 8; j++ {
			for _, char := range []byte(word) {
				if int(char) < 32 || int(char) > 126 {
					return "", fmt.Errorf("illegal character found in text: %s", string(char))
				}
				out.WriteString(data[int(char-32)*8+j])
			}
			out.WriteString("\n")
		}
	}
	return out.String(), nil
}

func ReadBanner(path string) ([]string, error) {
	panic("unimplemented")
}

func GetAscii(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open banner file: %w", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("cannot read banner file: %w", err)
	}
	return data, nil
}