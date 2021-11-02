package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// Get List of collections in a database.
func (srv *MongoRPCServer) ListCollections(ctx context.Context, in *proto.ListCollectionsRequest) (*proto.ListCollectionsResponse, error) {

	// Get Collections
	collections, err := srv.DB.Database(in.Database).ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	// Convert []string to []interface{}
	var arr []interface{}
	for _, collection := range collections {
		arr = append(arr, collection)
	}

	// Return Response Object
	return &proto.ListCollectionsResponse{
		Collections: &proto.ArrayValue{
			Values: EncodeArray(arr).Values,
		},
	}, nil
}

func (srv *MongoRPCServer) CreateCollection(ctx context.Context, in *proto.CreateCollectionRequest) (*proto.CreateCollectionResponse, error) {

	// Create Collection
	err := srv.DB.Database(in.Database).CreateCollection(ctx, in.Collection)
	if err != nil {
		return nil, err
	}

	// Return Response Object
	return &proto.CreateCollectionResponse{}, nil
}

func (srv *MongoRPCServer) RenameCollection(ctx context.Context, in *proto.RenameCollectionRequest) (*proto.RenameCollectionResponse, error) {

	// Rename Collection

	result, err := srv.RunDatabaseCommand(ctx, in.Database, bson.D{
		bson.E{
			Key:   "renameCollection",
			Value: in.Collection,
		},
		bson.E{
			Key:   "to",
			Value: in.Name,
		},
	})

	if err != nil {
		return nil, err
	}

	logrus.Debug(result)

	// Return Response Object
	return &proto.RenameCollectionResponse{}, nil
}

func (srv *MongoRPCServer) DeleteCollection(ctx context.Context, in *proto.DeleteCollectionRequest) (*proto.DeleteCollectionResponse, error) {

	// Delete Collection
	err := srv.DB.Database(in.Database).Collection(in.Collection).Drop(ctx)
	if err != nil {
		return nil, err
	}

	// Return Response Object
	return &proto.DeleteCollectionResponse{}, nil
}

func (srv *MongoRPCServer) CollectionStats(ctx context.Context, in *proto.CollectionStatsRequest) (*proto.CollectionStatsResponse, error) {

	// Get Collection Stats
	result, err := srv.RunDatabaseCommand(ctx, in.Database, bson.D{
		bson.E{
			Key:   "collStats",
			Value: in.Collection,
		},
	})

	if err != nil {
		return nil, err
	}

	var stats proto.CollectionStatsResponse
	err = result.Decode(&stats)
	if err != nil {
		return nil, err
	}

	// Return Response Object
	return &stats, nil
}
