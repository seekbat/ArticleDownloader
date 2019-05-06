package database

import (
	"context"
	"fmt"
	"github.com/seekbat/ArticleDownloader/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
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
	client *mongo.Client
	ctx    context.Context
}

func NewDatabase(connectString string, ctx context.Context) *Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectString))
	abortIfError(err)

	return &Database{client, ctx}
}

func (d *Database) AddLinksToDb(linklist models.LinkList) {

	err := d.client.Ping(d.ctx, nil) // Check the connection
	abortIfError(err)
	collection := d.client.Database("ArticleDownloader").Collection(linklist.SiteName)
	for _, link := range linklist.Links {
		_, err := collection.InsertOne(d.ctx, link)
		logError(err)
	}
}

func abortIfError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func logError(e error) {
	if e != nil {
		if e == context.DeadlineExceeded {
			fmt.Println("asd")
		} else if _, ok := e.(*mongo.WriteException); ok { //check if it is of the type mongo.WriteException
			fmt.Println("dsa")
		} else {
			fmt.Println(e)
			fmt.Println(reflect.TypeOf(e))
		}
	}
}
