package main

import (
	"context"
	"fmt"
	"github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) DeleteBlog(ctx context.Context, in *proto.BlogId) (*emptypb.Empty,
	error) {
	log.Printf("--- Client called DeleteBlog ---")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot parse ID")
	}

	res, err := collection.DeleteOne(ctx,
		bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Was not found!")
	}

	return &emptypb.Empty{}, nil
}
