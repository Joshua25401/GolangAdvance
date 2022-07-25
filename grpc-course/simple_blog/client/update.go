package main

import (
	"context"
	"github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"log"
)

func updateBlog(c proto.BlogServiceClient, id string) {
	log.Println("--- updateBlog was invoked! ---")

	newBlog := &proto.Blog{
		Id:       id,
		AuthorId: "Chesya!",
		Title:    "A New Title from Chesya :)",
		Content:  "Content in chesya blog!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating : %v\n", err)
	}

	log.Println("Blog was updated!")
}
