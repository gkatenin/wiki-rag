package main

import (
	"encoding/json"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetWikipediaTitle returns an article title based on QID.
func GetWikipediaTitle(qid string, lang string) (string, error) {
	const baseURL = "https://www.wikidata.org/w/api.php"

	params := url.Values{}
	params.Set("action", "wbgetentities")
	params.Set("ids", qid)
	params.Set("format", "json")
	params.Set("props", "sitelinks")
	params.Set("sitefilter", lang+"wiki")

	fullURL := baseURL + "?" + params.Encode()
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]any
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	entities := result["entities"].(map[string]any)
	entity := entities[qid].(map[string]any)
	sitelinks := entity["sitelinks"].(map[string]any)
	site := sitelinks[lang+"wiki"].(map[string]any)
	title := site["title"].(string)

	return title, nil
}

// GetWikipediaArticle returns an article based on its title and language.
func GetWikipediaArticle(title string, lang string) (string, error) {
	const apiURL = "https://%s.wikipedia.org/w/api.php?%s"

	params := url.Values{}
	params.Set("action", "query")
	params.Set("format", "json")
	params.Set("prop", "extracts")
	params.Set("explaintext", "1")
	params.Set("titles", title)

	fullURL := fmt.Sprintf(apiURL, lang, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]any
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	pages := result["query"].(map[string]any)["pages"].(map[string]any)
	for _, v := range pages {
		page := v.(map[string]any)
		return page["extract"].(string), nil
	}

	return "", fmt.Errorf("article not found")
}

// Data structure to unmarshal json response.
type SearchResponse struct {
	Query struct {
		Search []struct {
			Title   string `json:"title"`
			Snippet string `json:"snippet"`
		} `json:"search"`
	} `json:"query"`
}

// SearchWikipedia looks for a string given language and number of results.
func SearchWikipedia(query string, lang string, limit int) ([]string, error) {
	const apiURL = "https://%s.wikipedia.org/w/api.php?%s"

	params := url.Values{}
	params.Set("action", "query")
	params.Set("list", "search")
	params.Set("srsearch", query)
	params.Set("format", "json")
	params.Set("srlimit", strconv.Itoa(limit))

	fullURL := fmt.Sprintf(apiURL, lang, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result SearchResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	titles := []string{}
	for _, item := range result.Query.Search {
		titles = append(titles, item.Title)
	}
	return titles, nil
}

