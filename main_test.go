package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var sentences = []string{
	"今日は良い天気だった。",
	"すごくだるい",
	"速達で追跡つけて送って欲しいと頼んだんだけど、明日に届けばいいなら要らないと思った。",
}

func TestTokenize(t *testing.T) {
	fmt.Print(tokenize(sentences))
}

func TestMakeSentence(t *testing.T) {
	tokens := processedToken{rest: []string{"今日", "は", "良い", "天気", "だっ", "た", "。", "すごく", "だるい", "速達", "で", "追跡", "つけ", "て", "送っ", "て", "欲しい", "と", "頼ん", "だ", "ん", "だ", "けど", "、", "明日", "に", "届け", "ば", "いい", "なら", "要ら", "ない", "と", "思っ", "た"}}

	fmt.Print(tokens.makeSentence(50))
}

func TestCreateFile(t *testing.T) {
	data := "aaaaaaaa"

	file, err := createFile(strings.NewReader(data))
	if err != nil {
		t.Errorf("create file fail %s", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()
}
