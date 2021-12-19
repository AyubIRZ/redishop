package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ayubirz/redishop/entity"
	"github.com/ayubirz/redishop/util"
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
)

type cartRedisRepo struct{}

var carts []entity.Cart
var cartRedisPrefix = "cart_"

// NewProductMysqlRepository creates a new instance of mysql provider repo for product entity
func NewCartRedisRepository() CartRepository {
	return &cartRedisRepo{}
}

// Find looks up and returns a requested product by its ID
func (*cartRedisRepo) Find(id int64) (entity.Cart, error) {
	idStr := strconv.FormatInt(id, 10)
	redisCartId := cartRedisPrefix + idStr

	client := util.GetRedisClient()

	cartJson, err := client.Get(redisCartId).Result()
	if err == redis.Nil {
		return entity.Cart{}, &CartNotFoundError{}
	} else if err != nil {
		return entity.Cart{}, errors.New(fmt.Sprintf("error fetching cart form redis: %v", err))
	}

	cart := entity.Cart{}

	err = json.Unmarshal([]byte(cartJson), &cart)
	if err != nil {
		return entity.Cart{}, errors.New(fmt.Sprintf("error decoding cart with id %d form redis: %v", id, err))
	}

	return cart, nil
}

// Save
func (r *cartRedisRepo) Save(cart entity.Cart) (*entity.Cart, error) {
	finalID := cart.ID
	if cart.ID == 0 {
		finalID = rand.Int63()
	}

	idStr := strconv.FormatInt(finalID, 10)
	cart.ID = finalID

	redisCartId := cartRedisPrefix + idStr

	client := util.GetRedisClient()

	encodedCart, err := json.Marshal(cart)
	if err != nil {
		return &entity.Cart{}, errors.New(fmt.Sprintf("error encoding cart to save into redis: %v", err))
	}

	err = client.Set(redisCartId, string(encodedCart), 0).Err()
	if err != nil {
		return &entity.Cart{}, errors.New(fmt.Sprintf("error saving cart into redis: %v", err))
	}

	return &cart, nil
}
