package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	Traffic  string `xml:"approx_traffic"`
	NewsItem []News `xml:"news_item"`
}

type News struct {
	HeadLine     string `xml:"news_item_title"`
	HeadLineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS
	data := readGoogleTrend()
	xml.Unmarshal(data, &r)

}

func readGoogleTrend() []byte {
	resp := getGoogleTrend()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}
	return data
}

func getGoogleTrend() *http.Response {
	resp, err := http.Get("https://trends.google.com.tw/trends/trendingsearches/daily/rss?geo=TW")
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}
	return resp
}
