package repository

import (
	"errors"
	"fmt"
	"github.com/ayubirz/redishop/entity"
	"strings"
)

var products []entity.Product

type inMemRepo struct{}

// NewProductMysqlRepository creates a new instance of mysql provider repo for product entity
func NewProductInMemoryRepository() ProductRepository {
	return &inMemRepo{}
}

// Find looks up and returns a requested product by its ID
func (*inMemRepo) Find(id int64) (entity.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return entity.Product{}, errors.New(fmt.Sprintf("product with id %d not found", id))
}

// FindAll
func (*inMemRepo) FindAll() (*[]entity.Product, error) {
	if len(products) > 0 {
		return &products, nil
	}

	return &[]entity.Product{}, errors.New("no product found")
}

// SearchByTitle
func (*inMemRepo) SearchByTitle(term string) (*[]entity.Product, error) {
	var res []entity.Product

	for _, product := range products {
		if strings.Contains(product.Title, term) {
			res = append(res, product)
		}
	}

	return &res, nil
}

// Save
func (r *inMemRepo) Save(product entity.Product) (*entity.Product, error) {
	if _, err := r.Find(product.ID); err == nil {
		return &entity.Product{}, errors.New(fmt.Sprintf("product with id %d already exists", product.ID))
	}

	product.ID = int64(len(products) + 1)
	products = append(products, product)

	return &product, nil
}
