package main

import (
	"context"
	"github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"log"
)

func createBlog(c proto.BlogServiceClient) string {
	log.Println("--- Create Blog was Invoked ---")

	// Create instance blog
	blog := &proto.Blog{
		AuthorId: "Joshua",
		Title:    "Joshua First Blog",
		Content:  "Content of gRPC service in Golang!",
	}

	// Call gRPC service
	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatal("Unexpected Error :", err)
	}

	log.Printf("Blog has been created : %s\n", res.Id)
	return res.Id
}
