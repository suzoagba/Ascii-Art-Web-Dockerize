package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	FirstRune      = ' '
	LastRune       = '~'
	LinesPerLetter = 8
)

type banner map[rune][][]byte

// Create a new banner from a file
func NewBanner(bannerName string) (b banner, err error) {
	b = make(banner)

	// Read the banner file
	rawBanner, err := os.ReadFile("banners/" + bannerName + ".txt")
	if err != nil {
		return
	}

	// Convert the Windows line endings to Unix line endings
	rawBanner = bytes.ReplaceAll(rawBanner, []byte("\r\n"), []byte("\n"))

	// Trim the leading and trailing newlines
	rawBanner = bytes.Trim(rawBanner, "\n")

	// Split the banner into letters
	letters := bytes.Split(rawBanner, []byte("\n\n"))

	// Create a map of letters
	for i, letter := range letters {
		b[rune(FirstRune+i)] = bytes.Split(letter, []byte("\n"))
	}

	return
}

// Print a single line of text as ASCII art
func (b banner) PrintLine(text string) (asciiArt []byte, err error) {

	// Print the text, one line at a time
	for i := 0; i < LinesPerLetter; i++ {
		for _, c := range text {
			// Append the line to the ASCII art
			asciiArt = append(asciiArt, b[c][i]...)
		}
		asciiArt = append(asciiArt, '\n')
	}

	return
}

// Print multiple lines of text as ASCII art
func (b banner) PrintText(text string) (asciiArt []byte, err error) {

	// Check if the text is printable
	for _, c := range text {
		if c < FirstRune || c > LastRune {
			return nil, fmt.Errorf("invalid character: %q", c)
		}
	}

	// Split the text into lines
	lines := strings.Split(text, "\\n")

	// Print each line
	for i, line := range lines {
		if line == "" { // Print a blank line if the line is empty
			if i != 0 { // Except if it's the first line (weird task requirement)
				asciiArt = append(asciiArt, '\n')
			}
		} else { // Print the line using PrintLine
			s, err := b.PrintLine(line)
			if err != nil {
				return nil, err
			}
			asciiArt = append(asciiArt, s...)
		}
	}
	return
}