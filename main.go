package main

import (
	"fmt"
	"github.com/ayubirz/redishop/controller"
	"github.com/ayubirz/redishop/indexer"
	"github.com/ayubirz/redishop/repository"
	"github.com/ayubirz/redishop/router"
	"github.com/ayubirz/redishop/search"
	"github.com/ayubirz/redishop/service"
	"log"
	"net/http"
	"os"
)

var (
	httpRouter        router.Router                = router.NewMuxRouter()
	productRepository repository.ProductRepository = repository.NewProductMysqlRepository()
	productService    service.ProductService       = service.NewProductService(productRepository)
	productSearch     search.ProductSearch         = search.NewProductRedisSearch()
	productController controller.ProductController = controller.NewProductController(productService, productSearch, httpRouter)
	cartRepository    repository.CartRepository    = repository.NewCartRedisRepository()
	cartService       service.CartService          = service.NewCartService(cartRepository, )
	cartController    controller.CartController    = controller.NewCartController(cartService, httpRouter)
)

func main() {
	log.SetOutput(os.Stdout)

	index := indexer.NewProductIndexer(productService)
	index.IndexProducts()

	port := ":80"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/products", productController.GetAll)
	httpRouter.POST("/products", productController.Add)
	httpRouter.GET("/products/search/{term}", productController.SearchByTitle)
	httpRouter.GET("/carts/{id}", cartController.GetCart)
	httpRouter.POST("/carts/product", cartController.AddProductToCart)
	httpRouter.DELETE("/carts/product", cartController.RemoveProductFromCart)

	log.Fatal(httpRouter.SERVE(port))
}
