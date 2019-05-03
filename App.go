package main

import (
	"context"
	"fmt"
	"github.com/seekbat/ArticleDownloader/LinkScraper"
	"github.com/seekbat/ArticleDownloader/database
	"github.com/seekbat/ArticleDownloader/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"regexp"
	"strconv"
	"time"
)
var Sites []models.Site
var linklists []models.LinkList

func main() {
	min := models.Site{
		ID:        1,
		Name:      "20min",
		URL:       "https://www.20min.ch",
		LinkRegex: `\/([A-Za-z0-9-]{1,})([0-9]{1,}$)`,
		IDRegex: `([0-9]{1,}$)`,
	}
	db := database.NewDatabase("mongodb://localhost:27017", context.TODO())



	r, _ := regexp.Compile(min.LinkRegex)
	for _,site := range Sites {
		var articlelinks []models.ArticleLink
		articlelinks = linkscraper.LinkScraper(site.URL, site.LinkRegex, site.IDRegex)
		var linklist = models.LinkList{site.Name,articlelinks}
		linklists = append(linklists, linklist)
	}

	for _,linklist := range linklists  {
		db.AddLinksToDb(linklist)
	}


}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
