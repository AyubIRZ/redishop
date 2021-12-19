package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ayubirz/redishop/entity"
	"github.com/ayubirz/redishop/router"
	"github.com/ayubirz/redishop/search"
	"github.com/ayubirz/redishop/service"
	"log"
	"net/http"
)

var (
	productService service.ProductService
	productSearch search.ProductSearch
	productRouter  router.Router
)

// ProductController
type ProductController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Add(response http.ResponseWriter, request *http.Request)
	SearchByTitle(response http.ResponseWriter, request *http.Request)
}

// NewProductController
func NewProductController(service service.ProductService, search search.ProductSearch, router router.Router) ProductController {
	productService = service
	productSearch = search
	productRouter = router

	return &controller{}
}

// GetAll
func (*controller) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	products, err := productService.FindAll()
	if err != nil {
		log.Printf("Error getting the products: %v \n", err)
		response.WriteHeader(http.StatusInternalServerError)
		resp := ErrorResponse("Error getting the products")
		_ = json.NewEncoder(response).Encode(resp)

		return
	}

	if len(*products) == 0 {
		response.WriteHeader(http.StatusNotFound)
		resp := ErrorResponse("no product found")
		_ = json.NewEncoder(response).Encode(resp)

		return
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(products)
}

// Add
func (*controller) Add(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var product entity.Product

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(ErrorResponse("Error decoding product data"))
		return
	}

	ok, err1 := productService.Validate(product)
	if err1 != nil && ok == false {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(response).Encode(ErrorResponse(fmt.Sprintf("invalid product data")))

		return
	}

	result, err2 := productService.Create(product)
	if err2 != nil {
		log.Printf("Error saving the product information: %v \n", err2)
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(ErrorResponse("Error saving the product"))

		return
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(result)
}

// SearchByTitle
func (*controller) SearchByTitle(response http.ResponseWriter, request *http.Request) {
	params := productRouter.GetVars(request)
	searchTerm := params["term"]

	response.Header().Set("Content-Type", "application/json")
	products, err := productSearch.SearchByTitle(searchTerm)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		resp := ErrorResponse("Error getting the search results for products")
		_ = json.NewEncoder(response).Encode(resp)

		return
	}

	if len(*products) == 0 {
		response.WriteHeader(http.StatusNotFound)
		resp := ErrorResponse("no product found")
		_ = json.NewEncoder(response).Encode(resp)

		return
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(products)
}
