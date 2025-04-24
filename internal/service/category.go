package service

import (
	"context"
    "io"

	"github.com/silasstoffel/Go-gRPC/internal/pb"
	"github.com/silasstoffel/Go-gRPC/internal/database"
)


type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryInput) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(input.Name, input.Description)
    if err != nil {
        return nil, err
    }

    return &pb.Category{
		Id:  category.ID,
		Name: category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) GetCategories(ctx context.Context, input *pb.Blank) (*pb.GetCategoriesOutput, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.GetCategoriesOutput{Categories: pbCategories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, input *pb.GetCategoryInput) (*pb.Category, error) {
    category, err := c.CategoryDB.Find(input.Id)
    if err != nil {
        return nil, err
    }

    return &pb.Category{
        Id:          category.ID,
        Name:        category.Name,
        Description: category.Description,
    }, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
    categories := &pb.GetCategoriesOutput{
        Categories: []*pb.Category{},
    }
    for {
        input, err := stream.Recv()

        if err == io.EOF {
            return stream.SendAndClose(categories)
        }

        if err != nil {
            return err
        }

        category, err := c.CategoryDB.Create(input.Name, input.Description)
        if err != nil {
            return err
        }

        categories.Categories = append(categories.Categories, &pb.Category{
            Id:          category.ID,
            Name:        category.Name,
            Description: category.Description,
        })
    }
}
