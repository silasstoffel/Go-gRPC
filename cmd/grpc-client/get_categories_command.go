package main

import (
	"context"
	"log"
	"time"
	"github.com/silasstoffel/Go-gRPC/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCategoryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.GetCategories(ctx, &pb.Blank{})

	if err != nil {
		log.Fatalf("could not get categories: %v", err)
	}

	log.Println("Categories:")
	for _, category := range response.Categories {
		log.Printf("Id: %s, Name: %s, Description: %s", category.Id, category.Name, category.Description)
	}
}
