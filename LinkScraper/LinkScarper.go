package LinkScraper

import (
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func LinkScraper(url string, regex *regexp.Regexp) []string {
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

	var links []string

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		// See if the href attribute exists on the element
		href, exists := element.Attr("href")
		if exists {
			cleanLink, found := checkLink(href, regex)
			if found {
				links = appendIfNotExists(links, href)
			}
		}
	})
	return links
}

func checkLink(link string, regex *regexp.Regexp) (string, bool) {
	if len(link) < 1 {
		return "", false
	}
	if regex.MatchString(link) {
		return link, true
	}
	return "", false
}

/*

 */
func appendIfNotExists(strings []string, newString string) []string {
	exists := false
	for _, existingString := range strings {
		if existingString == newString {
			exists = true
		}
	}
	if !exists {
		strings = append(strings, newString)
	}
	return strings
}
