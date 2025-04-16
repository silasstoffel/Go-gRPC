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

	category := &pb.CreateCategoryInput{
		Name:        "Infrastructure",
		Description: "Infrastructure courses",
	}

	response, err := c.CreateCategory(ctx, category)
	if err != nil {
		log.Fatalf("could not create category: %v", err)
	}

	log.Printf("Category created. Id: %s, Name: %s, Description: %s", response.Id, response.Name, response.Description)
}
