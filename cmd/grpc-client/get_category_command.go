package main

import (
	"context"
	"log"
	"time"
	"github.com/silasstoffel/Go-gRPC/internal/pb"
	"google.golang.org/grpc"
)

func main() {
    id := "f2791b29-4a5b-435a-969f-724beb708b8b"
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCategoryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.GetCategory(ctx, &pb.GetCategoryInput{Id: id})

	if err != nil {
		log.Fatalf("Could not get category by id: %v Detail: %v", id, err)
	}

    log.Printf("Id: %s", response.Id)
    log.Printf("Name: %s", response.Name)
    log.Printf("Description: %s", response.Description)
}
