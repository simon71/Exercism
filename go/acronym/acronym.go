// Convert a phrase to its acronym.

// Techies love their TLA (Three Letter Acronyms)!

// Help generate some jargon by writing a program that converts a long name like Portable Network Graphics to its acronym (PNG).

package acronym

// package main

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	acronym := ""

	// Split the string into individual words
	for _, words := range strings.SplitAfter(s, " ") {

		// convert into a rune
		words := []rune(words)

		// Grab the first letter of each word
		letters := string(words[0])

		// convert to uppercase
		upperLetters := strings.ToUpper(letters)

		acronym += upperLetters
	}

	return acronym
}
