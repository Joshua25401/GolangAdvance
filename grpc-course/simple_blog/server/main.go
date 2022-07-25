package main

import (
	"context"
	pb "github.com/joshuaryandafres/golang/grpc-course/simple_blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = ":5051"

var collection *mongo.Collection
var connectionURI = "mongodb://joshua:capitalX123@localhost:27017"

type Server struct {
	pb.BlogServiceServer
}

func initDB() {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))

	if err != nil {
		log.Fatalln("Error making mongoClient", err)
	}

	err = mongoClient.Connect(context.Background())

	if err != nil {
		log.Fatalln("Unable to connect", err)
	}

	collection = mongoClient.Database("blogdb").Collection("blog")
}

func main() {

	// Initiate mongoDB
	initDB()

	// Initiate server
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Error when trying to listen", err)
	}
	log.Println("Success listening on", addr)

	server := grpc.NewServer()

	pb.RegisterBlogServiceServer(server, &Server{})

	if err = server.Serve(listen); err != nil {
		log.Fatalln("Errong when trying to make server", err)
	}
}
