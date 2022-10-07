package main

/*
go run main.go -s zh -st 今天天氣不錯 -t en

result: the weather is nice today
*/

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/shaiuwen0310/google-translate/cli"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "sourceLanguage[en]")
	flag.StringVar(&targetLang, "t", "fr", "sourceLanguage[fr]")
	flag.StringVar(&sourceText, "st", "", "文字翻譯")
}

func main() {
	// 解析輸入的參數
	flag.Parse()
	fmt.Println("參數: ", sourceLang, targetLang, sourceText)

	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	strCh := make(chan string)

	//wg add, adds a counter, done reduces by 1 and wait waits for it to hit 0
	wg.Add(1)

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	go cli.RequestTranslate(reqBody, strCh, &wg)

	processedStr := strings.ReplaceAll(<-strCh, "+", " ")
	// processedStr := <-strCh
	fmt.Printf("%s\n", processedStr)
	close(strCh)
	wg.Wait()
}
