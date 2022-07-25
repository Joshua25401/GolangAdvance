package main

import (
	"context"
	"fmt"
	"github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateBlog(ctx context.Context, in *proto.Blog) (*proto.BlogId,
	error) {
	log.Println("--- Client called CreateBlog ---")

	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error : %v\n", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID")
	}

	return &proto.BlogId{
		Id: oid.Hex(),
	}, nil
}
