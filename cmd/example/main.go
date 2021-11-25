package main

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	c := client.NewClient("localhost:9090")
	conn, err := c.Connect(
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()

	// Initilize database
	db := c.Database("sample_mflix")

	// Get Document By ID
	doc, err := db.Collection("movies").Document("573a1390f29313caabcd4135").Get(context.TODO())
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infoln(doc)

	// docs, err := db.Collection("movies").Documents().Limit(10).Skip(10).Sort("name", client.ASCENDING).Get(context.TODO())
	// docs, err := db.Collection("movies").Documents().EqualTo("title", "Gertie the Dinosaur").Get(context.TODO())
	docs, err := db.Collection("movies").Documents().Search("Batman").Get(context.TODO())
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infoln(docs)
}
