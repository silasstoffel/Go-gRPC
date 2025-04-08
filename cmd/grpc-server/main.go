package main

import (
	"database/sql"

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
	
	if _, error := net.Listen("tcp", ":50051"); error != nil {
		panic(error)
	}
}

func migrate(db *sql.DB) {
	// Create a table
	_, error := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if error != nil {
		panic(error)
	}

	// Create a table
	_, error = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if error != nil {
		panic(error)
	}
}