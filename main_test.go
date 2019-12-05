package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var sentences = []string{
	"今日は良い天気だった。",
	"すごくだるい",
	"速達で追跡つけて送って欲しいと頼んだんだけど、明日に届けばいいなら要らないと思った。",
}

func init() {
	myTwitter = "/"
}

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/twitter_ykpythemind_response.html")
	})

	return httptest.NewServer(mux)
}

func TestTokenize(t *testing.T) {
	fmt.Print(tokenize(sentences))
}

func TestMakeSentence(t *testing.T) {
	tokens := processedToken{rest: []string{"今日", "は", "良い", "天気", "だっ", "た", "。", "すごく", "だるい", "速達", "で", "追跡", "つけ", "て", "送っ", "て", "欲しい", "と", "頼ん", "だ", "ん", "だ", "けど", "、", "明日", "に", "届け", "ば", "いい", "なら", "要ら", "ない", "と", "思っ", "た"}}

	fmt.Print(tokens.makeSentence(50))
}

func TestTwitterScrap(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	tw := newTwitterScraper(ts.URL)
	arr, err := tw.Exec()
	if err != nil {
		t.Fatalf("exec err: %s", err)
	}

	if !strings.Contains(arr[0], "【RT希望】") {
		t.Error("fail to parse tweet?")
	}
}
