package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func (d *Database) AddLinksToDb() {

	err := d.client.Ping(d.ctx, nil) // Check the connection
	checkErr(err)

}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
