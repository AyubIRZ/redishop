package search

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/ayubirz/redishop/entity"
	"github.com/ayubirz/redishop/util"
	"log"
	"strconv"
)

type productRedisSearch struct {}

func NewProductRedisSearch() ProductSearch {
	return &productRedisSearch{}
}

func (*productRedisSearch) SearchByTitle(term string) (*[]entity.Product, error) {
	client := util.GetRedisSearchClient()

	var searchResult []entity.Product

	// Searching with limit and sorting
	docs, total, err := client.Search(redisearch.NewQuery(term))
	if err != nil {
		return &searchResult, err
	}
	log.Println(total)

	for _, doc := range docs {
		log.Println(doc.Payload)

		id, _ := strconv.ParseInt(doc.Id, 10, 64)

		pr := entity.Product{
			ID:          id,
			Title:       doc.Properties["title"].(string),
			Description: doc.Properties["description"].(string),
		}

		searchResult = append(searchResult, pr)
	}

	return &searchResult, nil
}