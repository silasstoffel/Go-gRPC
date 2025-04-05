package service

import "github.com/silasstoffel/Go-gRPC/internal/pb"
import "github.com/silasstoffel/Go-gRPC/database"

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(categoryDB *database.Category) *CategoryService {
    return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryInput) (*Category, error) {
	category, err := c.CategoryDB.Create(input.Name, input.Description)
    if err != nil {
        return nil, err
    }
	
    return &pb.Category{
		Id: category.ID, 
		Name: category.Name, 
		Description: category.Description
	}, nil
}