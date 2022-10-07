package cli

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

// https://translate.googleapis.com/translate_a/single?client=gtx&dt=t&sl=en&tl=zh-CN&q="test"
const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(b *RequestBody, str chan string, wg *sync.WaitGroup) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", translateUrl, nil)

	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", b.SourceLang)
	query.Add("tl", b.TargetLang)
	query.Add("dt", "t")
	query.Add("q", b.SourceText)

	// 將get後面的參數加到url中
	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("client Do: ", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "請求過多，晚點再試一次"
		wg.Done()
		return
	}

	parseJson, err := gabs.ParseJSONBuffer(res.Body)
	fmt.Println("從body中取出json: ", parseJson)
	if err != nil {
		log.Fatal("gabs ParseJSONBuffer: ", err)
	}

	nestOne, err := parseJson.ArrayElement(0)
	if err != nil {
		log.Fatal("nestOne: ", err)
	}

	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatal("nestTwo: ", err)
	}

	translatedStr, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatal("translatedStr: ", err)
	}

	str <- translatedStr.Data().(string)
	wg.Done()
}
