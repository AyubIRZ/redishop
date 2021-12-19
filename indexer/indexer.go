package indexer

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/ayubirz/redishop/service"
	"github.com/ayubirz/redishop/util"
	"log"
	"strconv"
)

type productIndexer struct {}

var (
	productService service.ProductService
)

func NewProductIndexer(service service.ProductService) *productIndexer {
	productService = service

	return &productIndexer{}
}

func (*productIndexer) IndexProducts() {
	c := util.GetRedisSearchClient()

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewTextField("description"))

	err := c.Drop()
	if err != nil {
		log.Fatal(err)
	}

	if err := c.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}

	products, err := productService.FindAll()
	if err != nil {
		log.Fatalf("Indexing products error: %v", err)
	}

	for _, product := range *products {
		doc := redisearch.NewDocument(strconv.FormatInt(product.ID, 10), 1.0)
		doc.Set("title", product.Title).Set("description", product.Description)

		// Index the document. The API accepts multiple documents at a time
		if err := c.Index([]redisearch.Document{doc}...); err != nil {
			log.Fatal(err)
		}
	}
}