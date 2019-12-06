package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ikawaha/kagome/tokenizer"
)

var myTwitter = "https://twitter.com/ykpythemind"
var myRepository = "https://github.com/ykpythemind/monkeydiary_web"
var myName = "ykpythemind"

type config struct {
	watchTwitterAccount string
	diaryRepository     string
	userName            string
}

func main() {
	config := newConfig()

	ts := newTwitterScraper(myTwitter)
	tweets, err := ts.Exec()
	if err != nil {
		log.Fatalf("twitter execution error %s", err)
	}

	rand.Seed(time.Now().Unix())

	token := tokenize(tweets[1:])                // 固定ツイートをスキップ
	res := token.makeSentence(rand.Intn(20) + 5) // 5...25

	log.Printf("generated diary. \n%s\n", res)

	r := strings.NewReader(res)
	g := newUploadGithubService(config)
	err = g.upload(r)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func newConfig() *config {
	watchTwitterAccount := os.Getenv("TWITTER_ACCOUNT_ID")
	if watchTwitterAccount == "" {
		watchTwitterAccount = myTwitter
	}

	diaryRepository := os.Getenv("TARGET_REPOSITORY_URL")
	if diaryRepository == "" {
		diaryRepository = myRepository
	}

	userName := os.Getenv("GITHUB_USERNAME")
	if userName == "" {
		userName = myName
	}

	return &config{
		watchTwitterAccount: watchTwitterAccount,
		diaryRepository:     diaryRepository,
		userName:            userName,
	}
}

type processedToken struct {
	first []string // TODO
	rest  []string
	last  []string // TODO
}

func (t processedToken) String() string {
	return fmt.Sprintf("first: %s\nrest : %s\nlast : %s\n", strings.Join(t.first, ","), strings.Join(t.rest, ","), strings.Join(t.last, ","))
}

func tokenize(strs []string) processedToken {
	var result processedToken

	for _, str := range strs {
		t := tokenizer.New()
		tokens := t.Tokenize(str)
		for _, token := range tokens {
			if token.Class == tokenizer.DUMMY {
				log.Printf("DUMMY: skip")
				continue
			}
			if token.Class == tokenizer.UNKNOWN {
				log.Printf("UNKNOWN: %s", token.String())
				continue
			}

			result.rest = append(result.rest, token.Surface)
		}
	}

	return result
}

func (t processedToken) makeSentence(tokenSize int) string {
	rand.Seed(time.Now().Unix())
	var samples []string

	for i := 0; i < 50; i++ {
		sampled := t.rest[rand.Intn(len(t.rest))]
		// log.Printf("token: %s", sampled)
		samples = append(samples, sampled)
	}

	return strings.Join(samples, "")
}
