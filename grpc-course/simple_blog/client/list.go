package main

import (
	"context"
	"fmt"
	"github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlog(c proto.BlogServiceClient) {
	fmt.Println("listBlog was invoked!")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while called service : %v\n", err)
	}

	for {
		res, err := stream.Recv()

		// Stream finished
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened : %v\n", err)
		}

		fmt.Println(res)
	}
}
