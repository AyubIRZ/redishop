package controller

import (
	"encoding/json"
	"errors"
	"github.com/ayubirz/redishop/entity"
	"github.com/ayubirz/redishop/repository"
	"github.com/ayubirz/redishop/router"
	"github.com/ayubirz/redishop/service"
	"log"
	"net/http"
	"strconv"
)

var (
	cartService service.CartService
	cartRouter router.Router
)

// CartController
type CartController interface {
	GetCart(response http.ResponseWriter, request *http.Request)
	AddProductToCart(response http.ResponseWriter, request *http.Request)
	RemoveProductFromCart(response http.ResponseWriter, request *http.Request)
}

// NewCartController
func NewCartController(service service.CartService, router router.Router) CartController {
	cartService = service
	cartRouter = router

	return &controller{}
}

// GetCart
func (*controller) GetCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := cartRouter.GetVars(request)
	cartId, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		resp := ErrorResponse("invalid cart ID")
		_ = json.NewEncoder(response).Encode(resp)
	}

	var notFoundError *repository.CartNotFoundError
	cart, err := cartService.GetCart(int64(cartId))
	if err != nil {
		if errors.As(err, &notFoundError) {
			response.WriteHeader(http.StatusNotFound)
			resp := ErrorResponse("Cart not found")
			_ = json.NewEncoder(response).Encode(resp)

			return
		}

		log.Printf("Error getting cart: %v \n", err)
		response.WriteHeader(http.StatusInternalServerError)
		resp := ErrorResponse("Error getting the requested cart")
		_ = json.NewEncoder(response).Encode(resp)

		return
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(cart)
}

// AddProductToCart
func (*controller) AddProductToCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var body = struct {
		ProductID int64 `json:"product_id"`
		Quantity int `json:"quantity"`
		CartID int64 `json:"cart_id"`
	}{}

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(ErrorResponse("Error decoding cart data"))

		return
	}

	if body.Quantity < 1 {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(response).Encode(ErrorResponse("quantity must be greater than 1"))

		return
	}

	product, err := productService.Find(body.ProductID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(response).Encode(ErrorResponse("given cart product not found"))

		return
	}

	cart, err := cartService.GetCart(body.CartID)
	if err != nil {
		cart = entity.Cart{}
	}

	cartItem := entity.CartItem{
		Product:  product,
		Quantity: body.Quantity,
	}

	cart, err = cartService.AddItemToCart(cartItem, cart)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(ErrorResponse(err.Error()))

		return
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(cart)
}

// RemoveProductFromCart
func (*controller) RemoveProductFromCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var body = struct {
		ProductID int64 `json:"product_id"`
		CartID int64 `json:"cart_id"`
	}{}

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(ErrorResponse("Error decoding cart data"))

		return
	}

	product, err := productService.Find(body.ProductID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(response).Encode(ErrorResponse("given cart product not found"))

		return
	}

	cart, err := cartService.GetCart(body.CartID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(response).Encode(ErrorResponse("cart not found"))

		return
	}

	cart, err = cartService.RemoveItemFromCart(product, cart)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(ErrorResponse(err.Error()))

		return
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(cart)
}
