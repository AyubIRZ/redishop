package repository

import (
	"github.com/ayubirz/redishop/entity"
)

type CartRepository interface {
	Find(id int64) (entity.Cart, error)
	Save(cart entity.Cart) (*entity.Cart, error)
}

type CartNotFoundError struct {}

func (e *CartNotFoundError) Error() string {
	return "cart not found"
}