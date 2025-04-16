package main

import (
	"net"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/silasstoffel/Go-gRPC/internal/database"
	"github.com/silasstoffel/Go-gRPC/internal/service"
	"github.com/silasstoffel/Go-gRPC/internal/pb"

)

func main() {
	db, error := sql.Open("sqlite3", "./database.sqlite")

	if error != nil {
		panic(error)
	}

	defer db.Close()

	migrate(db)

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func migrate(db *sql.DB) {
	// Create a table
	_, error := db.Exec("CREATE TABLE IF NOT EXISTS categories (id varchar(40) PRIMARY KEY, name varchar(100), description TEXT DEFAULT NULL)")
	if error != nil {
		panic(error)
	}
}
