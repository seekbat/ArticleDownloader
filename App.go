package main

import (
	"fmt"
	"github.com/seekbat/ArticleDownloader/LinkScraper"
	"regexp"
)

type Site struct {
	ID    int    `json:"id"`    //ID for Internal Use
	Name  string `json:"name"`  //Name of the Site
	URL   string `json:"url"`   //to the Site
	Regex string `json:"regex"` //to find te Links
}

func main() {

	min := Site{
		ID:    1,
		Name:  "20min",
		URL:   "https://www.20min.ch",
		Regex: `\/([A-Za-z0-9-]{1,})([0-9]{1,}$)`,
	}

	var links []string
	r, _ := regexp.Compile(min.Regex)
	links = LinkScraper.LinkScraper(min.URL, r)

	for _, link := range links {
		fmt.Println(link)
	}

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
