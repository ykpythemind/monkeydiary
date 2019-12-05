package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type twitterScraper struct {
	url string
}

func newTwitterScraper(url string) *twitterScraper {
	return &twitterScraper{url: url}
}

func (t twitterScraper) Exec() (body string, err error) {
	res, err := http.Get(t.url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	log.Print(doc.Text())
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Printf("article: %s\n", s.Text())
	})

	return body, nil
}
