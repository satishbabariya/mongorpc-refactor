package main

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	client := client.NewClient("localhost:9090")
	conn, err := client.Connect(
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()

	// Initilize database
	db := client.Database("sample_mflix")

	// Get Document By ID
	doc, err := db.Collection("movies").Document("573a1390f29313caabcd4135").Get(context.TODO())
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infoln(doc)
}
