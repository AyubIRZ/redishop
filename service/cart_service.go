package service

import (
	"errors"
	"github.com/ayubirz/redishop/entity"
	"github.com/ayubirz/redishop/repository"
)

type CartService interface {
	GetCart(id int64) (entity.Cart, error)
	AddItemToCart(cartItem entity.CartItem, cart entity.Cart) (entity.Cart, error)
	RemoveItemFromCart(product entity.Product, cart entity.Cart) (entity.Cart, error)
}

type cartService struct{}

var (
	cartRepo repository.CartRepository
)

func NewCartService(repository repository.CartRepository) CartService {
	cartRepo = repository

	return &cartService{}
}

func (*cartService) GetCart(id int64) (entity.Cart, error) {
	cart, err := cartRepo.Find(id)
	if err != nil {
		return entity.Cart{}, err
	}

	return cart, nil
}

func (s *cartService) AddItemToCart(cartItem entity.CartItem, cart entity.Cart) (entity.Cart, error) {
	if cart.ID != 0 {
		for _, item := range cart.Items {
			if item.Product.ID == cartItem.Product.ID {
				return entity.Cart{}, errors.New("product already exists in the cart")
			}
		}
	}

	cart.Items = append(cart.Items, cartItem)

	finalCart, err := cartRepo.Save(cart)
	if err != nil {
		return entity.Cart{}, err
	}

	return *finalCart, nil
}

func (s *cartService) RemoveItemFromCart(product entity.Product, cart entity.Cart) (entity.Cart, error) {
	var newItems []entity.CartItem
	for _, item := range cart.Items {
		if item.Product.ID != product.ID {
			newItems = append(newItems, item)
		}
	}

	cart.Items = newItems

	finalCart, err := cartRepo.Save(cart)
	if err != nil {
		return entity.Cart{}, err
	}

	return *finalCart, nil
}
