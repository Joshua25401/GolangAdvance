package main

import (
	"context"
	"fmt"
	"github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) ListBlogs(in *emptypb.Empty,
	stream proto.BlogService_ListBlogsServer) error {
	log.Println("--- Client called ListBlog ---")

	cur, err := collection.Find(context.Background(),
		primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error : %v\n", err))
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data %v\n", err))
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			"Unknown internal error")
	}
	return nil
}
