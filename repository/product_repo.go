package repository

import (
	"github.com/ayubirz/redishop/entity"
)

type ProductRepository interface {
	Find(id int64) (entity.Product, error)
	FindAll() (*[]entity.Product, error)
	SearchByTitle(term string) (*[]entity.Product, error)
	Save(product entity.Product) (*entity.Product, error)
}
