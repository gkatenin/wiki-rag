package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kelindar/search"
)

func buildIndex(v *search.Vectorizer, wt []string, lang string) *search.Index[string] {
	// We split each document into 1600 char chunks with an overlap of 100 chars.
	const (
		chunkLen  = 1600
		chunkOver = 100
	)
	index := search.NewIndex[string]()

	for _, article := range wt {
		text, _ := GetWikipediaArticle(article, lang)
		chunks := ChunkText(CleanText(text), chunkLen, chunkOver)

		for _, ch := range chunks {
			embedding, _ := v.EmbedText(ch)
			index.Add(embedding, ch)
		}
	}
	return index
}

func main() {
	const tmpl = `The context is given by the numbered text fragments below. Based on this context,
tell me about %q, using the language %q. Context:
`
	q := flag.String("query", "", "Wikipedia query")
	lang := flag.String("lang", "en", "Wikipedia language, e.g. 'en', 'ru'")
	flag.Parse()

	if *q == "" {
		fmt.Println("Specify the --query option to search")
		flag.Usage()
		os.Exit(1)
	}
	query := CleanText(*q)

	m, err := search.NewVectorizer("models/paraphrase-multilingual-MiniLM-L12-118M-v2-Q8_0.gguf", 0)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	titles, err := SearchWikipedia(query, *lang, 8)
	if err != nil {
		fmt.Println("Search returned error:", err, "generating without context..")
	}

	idx := buildIndex(m, titles, *lang)
	// Embed the query
	qe, _ := m.EmbedText(query)

	// Perform search in the vector database to find top 16 chunks.
	results := idx.Search(qe, 16)

	/* Re-rank document using BM25 relevance of documents to a given search query.
	ranked := rankDocuments(CleanText(*query), results)
	for i, idx := range ranked {
		if idx < 0 {
			continue
		}
		fmt.Printf("#%d (score rank): doc[%d] = %q\n", i+1, idx, results[idx].Value)
	}*/

	// Let's build the prompt.
	var prompt strings.Builder
	fmt.Fprintf(&prompt, tmpl, query, *lang)

	for i, ranked := range results {
		fmt.Fprintf(&prompt, "%d. %s\n", i+1, ranked.Value)
	}

	answer, err := ReaderLLM(prompt.String(), 2048)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}
