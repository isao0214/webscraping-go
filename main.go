package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Target URL
	url := "https://github.com/golang/go/tags"

	// Fetch HTML Document
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Panic(err)
	}

	// Target Element URL
	selector := "body > div.application-main > div > main > div.container-xl.clearfix.new-discussion-timeline.px-3.px-md-4.px-lg-5 > div > div.Box > div:nth-child(2) > div > div > div.d-flex > h4 > a"

	// Get Target Element in HTML Document
	elm := doc.Find(selector)

	// Get Text of Element
	txt := elm.Text()

	// Trim Text
	trimTxt := strings.TrimSpace(txt)

	// Display Text
	log.Println(trimTxt)
}
