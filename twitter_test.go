package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/twitter_ykpythemind_response.html")
	})

	return httptest.NewServer(mux)
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
