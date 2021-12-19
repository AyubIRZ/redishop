package service

import (
	"github.com/ayubirz/redishop/entity"
	"github.com/ayubirz/redishop/repository"
	"errors"
)

type productService struct{}

type ProductService interface {
	Validate(product entity.Product) (bool, error)
	Create(product entity.Product) (*entity.Product, error)
	FindAll() (*[]entity.Product, error)
	Find(id int64) (entity.Product, error)
	SearchByTitle(term string) (*[]entity.Product, error)
}

var (
	productRepo repository.ProductRepository
)

func NewProductService(repository repository.ProductRepository) ProductService {
	productRepo = repository
	return &productService{}
}

func (*productService) Validate(product entity.Product) (bool, error) {
	if product.Title == "" {
		return false, errors.New("The product title is empty")
	}

	return true, nil
}

func (*productService) Create(product entity.Product) (*entity.Product, error) {
	return productRepo.Save(product)
}

func (*productService) FindAll() (*[]entity.Product, error) {
	return productRepo.FindAll()
}

func (*productService) Find(id int64) (entity.Product, error) {
	return productRepo.Find(id)
}

func (*productService) SearchByTitle(term string) (*[]entity.Product, error) {
	return productRepo.SearchByTitle(term)
}
