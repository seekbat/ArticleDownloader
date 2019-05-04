package linkscraper

import (
	"fmt"
	"github.com/seekbat/ArticleDownloader/models"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func LinkScraper(url string, regex string, idregex string) []models.ArticleLink {

	// Compiling Regexes
	r, _ := regexp.Compile(regex)
	rid, _ := regexp.Compile(idregex)

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
	var linklist []models.ArticleLink

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		// See if the href attribute exists on the element
		href, exists := element.Attr("href")
		if exists {
			found := checkLink(href, r)
			if found {
				links = appendIfNotExists(links, href)
			}
		}

		for _, link := range links {
			id, err := strconv.Atoi(rid.FindString(link))
			checkErr(err)
			var article = models.ArticleLink{id, link, time.Now().Unix()}
			linklist = append(linklist, article)
		}
	})

	return linklist
}

func checkLink(link string, regex *regexp.Regexp) bool {
	if !(len(link) < 1) {
		return regex.MatchString(link)
	}

	return false
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

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
