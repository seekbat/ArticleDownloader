package main

import (
	"context"
	"fmt"
	"github.com/seekbat/ArticleDownloader/LinkScraper"
	"github.com/seekbat/ArticleDownloader/database
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"regexp"
	"strconv"
	"time"
)


func main() {
	min := Site{
		ID:        1,
		Name:      "20min",
		URL:       "https://www.20min.ch",
		LinkRegex: `\/([A-Za-z0-9-]{1,})([0-9]{1,}$)`,
		IDRegex: `([0-9]{1,}$)`,
	}

	var links []string
	r, _ := regexp.Compile(min.LinkRegex)
	links = LinkScraper.LinkScraper(min.URL, r)



	collection := client.Database("ArticleDownloader").Collection("20min")

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
