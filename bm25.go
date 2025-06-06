package main

import (
	"math"
	"strings"

	"github.com/kelindar/search"
)

type BM25 struct {
	docs         [][]string
	avgdl        float64
	docLens      []int
	docFreq      map[string]int
	k1, b, delta float64
	docCount     int
}

func NewBM25(tokenized [][]string, k1, b, delta float64) *BM25 {
	docFreq := make(map[string]int)
	totalLen := 0

	for _, words := range tokenized {
		totalLen += len(words)

		seen := make(map[string]bool)
		for _, w := range words {
			if !seen[w] {
				docFreq[w]++
				seen[w] = true
			}
		}
	}

	avgdl := float64(totalLen) / float64(len(tokenized))
	docLens := make([]int, len(tokenized))
	for i, d := range tokenized {
		docLens[i] = len(d)
	}

	return &BM25{
		docs:     tokenized,
		avgdl:    avgdl,
		docLens:  docLens,
		docFreq:  docFreq,
		k1:       k1,
		b:        b,
		delta:    delta,
		docCount: len(tokenized),
	}
}

func (bm *BM25) IDF(term string) float64 {
	df := float64(bm.docFreq[term])
	return math.Log((float64(bm.docCount)-df+0.5)/(df+0.5) + 1)
}

func (bm *BM25) Score(query string, docIdx int) float64 {
	score := 0.0
	qTerms := tokenize(query)
	tf := termFreq(bm.docs[docIdx])
	docLen := float64(bm.docLens[docIdx])

	for _, term := range qTerms {
		f := float64(tf[term])
		if f == 0 {
			continue
		}
		idf := bm.IDF(term)
		numerator := f * (bm.k1 + 1)
		denominator := f + bm.k1*(1-bm.b+bm.b*docLen/bm.avgdl)
		score += idf * (numerator/denominator + bm.delta)
	}
	return score
}

func termFreq(tokens []string) map[string]int {
	tf := make(map[string]int)
	for _, t := range tokens {
		tf[t]++
	}
	return tf
}

// Simple tokenizeer is splitting by whitespace.
func tokenize(text string) []string {
	return strings.Fields(text)
}

func rankDocuments(query string, documents []search.Result[string]) top3 {
	tokenizedDocs := make([][]string, len(documents))
	for i, doc := range documents {
		tokenizedDocs[i] = tokenize(doc.Value)
	}
	bm := NewBM25(tokenizedDocs, 1.9, 0.3, 1.9)

	var results []float64
	for i := range tokenizedDocs {
		s := bm.Score(query, i)
		results = append(results, s)
	}
	return TopThreeIndices(results)
}

/*
func main() {
	docs := []string{
		"Golang is great for concurrency and performance",
		"Python has a great ecosystem",
		"Concurrency is handled using goroutines in Go",
		"JavaScript runs in a single-threaded environment",
	}
	query := "golang concurrency"

	bm := NewBM25(docs, 1.5, 0.75, 1.0)

	type result struct {
		idx   int
		score float64
	}
	var results []result

	for i := range docs {
		s := bm.Score(query, i)
		results = append(results, result{i, s})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	for i, r := range results {
		fmt.Printf("#%d: doc[%d] (score %.4f): %q\n", i+1, r.idx, r.score, docs[r.idx])
	}
}*/
