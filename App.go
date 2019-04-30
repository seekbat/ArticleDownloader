package main

import (
	"context"
	"fmt"
	"github.com/seekbat/ArticleDownloader/LinkScraper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"regexp"
	"strconv"
	"time"
)

type Site struct {
	ID    int    `json:"_id"`   //ID for Internal Use
	Name  string `json:"name"`  //Name of the Site
	URL   string `json:"url"`   //to the Site
	Regex string `json:"regex"` //to find te Links
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
		ID:    1,
		Name:  "20min",
		URL:   "https://www.20min.ch",
		Regex: `\/([A-Za-z0-9-]{1,})([0-9]{1,}$)`,
	}

	var links []string
	r, _ := regexp.Compile(min.Regex)
	links = LinkScraper.LinkScraper(min.URL, r)

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	fmt.Println("Connected to MongoDB!")

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
