package database

import (
	"context"
	"fmt"
	"github.com/seekbat/ArticleDownloader"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"regexp"
	"strconv"
	"time"
)

/*
 * ==========================================
 *  Title:
 *  Project:   ArticleDownloader
 * Author:    Olivier Eggimann
 * Date:      29.Apr.2019
 * ==========================================
 */

type Database struct {
	op     *options.ClientOptions
	client *mongo.Client
	ctx    context.Context
}

func NewDatabase(op *options.ClientOptions, ctx context.Context) *Database {
	client, err := mongo.Connect(ctx, op)
	checkErr(err)

	return &Database{op, client, ctx}
}

func (d *Database) AddLinksToDb(linklist Link, regexid string) {
	rid, _ := regexp.Compile(regexid)
	err := d.client.Ping(d.ctx, nil) // Check the connection
	checkErr(err)
	for _, link := range links {
		fmt.Println(link)
		id, err := strconv.Atoi(rid.FindString(link))
		var link = ArticleDownloader.ArticleLink{id, link, time.Now().Unix()}
		_, err = collection.InsertOne(context.TODO(), link)
		if err != nil {
			fmt.Print(err)
		}
	}

}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
