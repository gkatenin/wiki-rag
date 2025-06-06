package main

import "testing"

var cleantextTests = []struct {
	input    string
	expected string
}{
	{" A ", "a"},
	{"", ""},
	{"     ", ""},
	{"   Я\t\n ", "я"},
	{"\n\t", ""},
	{"\f Ein übergroßer   FUẞGÄNGERÜBERGANG \r\t \n", "ein übergroßer fußgängerübergang"},
	{"Hello World", "hello world"},
	{"   Lots   of   spaces   ", "lots of spaces"},
	{"Tabs\tand\nnewlines\n", "tabs and newlines"},
	{"MIXED Case Input", "mixed case input"},
	{"\n\n\t  Leading and trailing\t \n", "leading and trailing"},
	{"\rMultiple   \t  \n types \n of   \t whitespace", "multiple types of whitespace"},
	{"   Hello    \n   World\t!", "hello world !"},
}

func TestCleanText(t *testing.T) {
	for _, test := range cleantextTests {
		result := CleanText(test.input)
		if result != test.expected {
			t.Errorf("CleanText(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func BenchmarkCleanText(b *testing.B) {
	for b.Loop() {
		for _, test := range cleantextTests {
			CleanText(test.input)
		}
	}
}

func TestTopThreeIndices(t *testing.T) {
	tests := []struct {
		input    []float64
		expected top3 // Expected indices of top values (any order)
	}{
		{[]float64{3.2, 5.5, 1.1, 7.8, 4.6, 6.3}, top3{3, 5, 1}},
		{[]float64{1.0, 2.0, 3.0}, top3{2, 1, 0}},
		{[]float64{9.0, 1.0}, top3{0, 1, -1}},
		{[]float64{2.2}, top3{0, -1, -1}},
		{[]float64{}, top3{-1, -1, -1}},
		{[]float64{-5.0, -1.0, -3.0, -2.0}, top3{1, 3, 2}},
		{[]float64{5.0, 5.0, 5.0, 1.0}, top3{0, 1, 2}}, // multiple duplicates
	}

	for _, test := range tests {
		result := TopThreeIndices(test.input)
		if result != test.expected {
			t.Errorf("TopThreeIndices: for input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
