package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"za": "https://google.co.za/search?q=",
}

type searchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildGoogleUrls(searchTerm, countryCode, languageCode,pages,count int)([]string,error){
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if goggleBase, found := googleDomains[countryCode];found{
		for i := 0, i <pages;i++{
			start := i*count
			scrapeUrl := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", goggleBase,searchTerm, count, languageCode, start)
		}
	}else{
		err := fmt.Errorf("Country %s is not currently supported", countryCode)
		return nil,err
	}
	return  toScrape,nil
}

func GoogleScrape(searchTerm, countryCode,languageCode string ,pages,count) ([]searchResult, err) {
	result := []searchResult{}
	resultCounter := 0
	googlePages,err:= buildGoogleUrls(searchTerm, countryCode,languageCode string ,pages,count)
	if err != nil{
		return nil,err
	}
}

func main() {

	res, err := GoogleScrape("Safed Jhnda", "com","en",1,30)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
