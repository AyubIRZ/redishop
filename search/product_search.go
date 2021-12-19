package search

import "github.com/ayubirz/redishop/entity"

type ProductSearch interface {
	SearchByTitle(term string) (*[]entity.Product, error)
}
