package main

import (
	"github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = ":5051"

func main() {
	// Create a connection
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("Error creating a connection!", err)
	}

	defer conn.Close()

	c := proto.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	//readBlog(c, "Non Existing ID")
	updateBlog(c, id)
	createBlog(c)
	listBlog(c)
	deleteBlog(c, id)
	readBlog(c, id)
}
