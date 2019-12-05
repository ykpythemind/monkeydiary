package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type twitterScraper struct {
	url string
}

func newTwitterScraper(url string) *twitterScraper {
	return &twitterScraper{url: url}
}

func (t twitterScraper) Exec() (tweets []string, err error) {
	res, err := http.Get(t.url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return []string{}, fmt.Errorf("status code error: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	var arr []string

	doc.Find(".js-tweet-text-container").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		arr = append(arr, strings.TrimSpace(t))
	})

	return arr, nil
}
