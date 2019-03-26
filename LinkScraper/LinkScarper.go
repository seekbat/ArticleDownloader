package LinkScraper

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func LinkScraper(url string, regex regexp.Regexp) []string {
	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		// See if the href attribute exists on the element
		href, exists := element.Attr("href")
		if exists {
			fmt.Println(href)
		}
	})
}
