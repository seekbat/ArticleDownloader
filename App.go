package main

import (
	"context"
	"github.com/seekbat/ArticleDownloader/LinkScraper"
	"github.com/seekbat/ArticleDownloader/database"
	"github.com/seekbat/ArticleDownloader/models"
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
		IDRegex:   `([0-9]{1,}$)`,
	}
	Sites = append(Sites, min)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db := database.NewDatabase("mongodb://localhost:27017", ctx)

	for _, site := range Sites {
		var articlelinks []models.ArticleLink
		articlelinks = linkscraper.LinkScraper(site.URL, site.LinkRegex, site.IDRegex)
		var linklist = models.LinkList{site.Name, articlelinks}
		linklists = append(linklists, linklist)
	}

	for _, linklist := range linklists {
		db.AddLinksToDb(linklist)
	}

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
