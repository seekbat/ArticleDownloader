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

type Site struct {
	ID        int    `json:"_id"`       //ID for Internal Use
	Name      string `json:"name"`      //Name of the Site
	URL       string `json:"url"`       //to the Site
	LinkRegex string `json:"linkregex"` //to find te Links
	IDRegex string `json:"idregex"` //to extract the ID
}

type LinkList struct {
	SiteName string        `json:"site_name"`
	Links    []ArticleLink `json:"links"`
}
type ArticleLink struct {
	ArticleId int    `json:"article_id" bson:"_id"`
	URL       string `json:"url"        bson:"url"`
	Timestamp int64  `json:"timestamp"  bson:"timestamp"`
}

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



	rid, _ := regexp.Compile(`([0-9]{1,}$)`)
	collection := client.Database("ArticleDownloader").Collection("20min")
	for _, link := range links {
		fmt.Println(link)
		id, err := strconv.Atoi(rid.FindString(link))
		var link = ArticleLink{id, link, time.Now().Unix()}
		_, err = collection.InsertOne(context.TODO(), link)
		if err != nil {
			fmt.Print(err)
		}
	}

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
