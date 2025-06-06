package main

import (
	"math"
	"strings"
	"unicode"
)

// CleanText replaces sequences of spaces, tabs, etc. with a single space.
func CleanText(line string) string {
	var b strings.Builder
	inSpace := false

	for _, r := range strings.TrimSpace(line) {
		if unicode.IsSpace(r) {
			if !inSpace {
				b.WriteRune(' ')
				inSpace = true
			}
		} else {
			b.WriteRune(unicode.ToLower(r))
			inSpace = false
		}
	}
	return b.String()
}

// ChunkText splits the input text into chunks of maxLength characters
// with an overlap.
func ChunkText(text string, maxLength, overlap int) []string {
	var chunks []string
	var start int

	runes := []rune(text)

	for start < len(runes) {
		end := min(start+maxLength, len(runes))

		chunks = append(chunks, string(runes[start:end]))
		start += maxLength - overlap
	}
	return chunks
}

type top3 [3]int

// TopThreeIndices returns indices of top 3 elements in a slice in one pass, the
// idea being iterating through the array, comparing each element with the current 
// three largest values, shifting down as necessary.
func TopThreeIndices(nums []float64) top3 {
	topIndices := top3{-1, -1, -1}
	topValues := []float64{math.Inf(-1), math.Inf(-1), math.Inf(-1)}

	for i, num := range nums {
		switch {
		case num > topValues[0]:
			topValues[2], topIndices[2] = topValues[1], topIndices[1]
			topValues[1], topIndices[1] = topValues[0], topIndices[0]
			topValues[0], topIndices[0] = num, i
		case num > topValues[1]:
			topValues[2], topIndices[2] = topValues[1], topIndices[1]
			topValues[1], topIndices[1] = num, i
		case num > topValues[2]:
			topValues[2], topIndices[2] = num, i
		}
	}
	return topIndices
}

