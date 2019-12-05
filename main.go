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

func main() {
	fmt.Printf("hoge")
	os.Exit(0)
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
		log.Printf("token: %s", sampled)
		samples = append(samples, sampled)
	}

	return strings.Join(samples, "")
}
