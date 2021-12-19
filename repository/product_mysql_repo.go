package repository

import (
	"errors"
	"fmt"
	"github.com/ayubirz/redishop/entity"
	"github.com/ayubirz/redishop/util"
)

type mysqlRepo struct{}

// NewProductMysqlRepository creates a new instance of mysql provider repo for product entity
func NewProductMysqlRepository() ProductRepository {
	return &mysqlRepo{}
}

// Find looks up and returns a requested product by its ID
func (*mysqlRepo) Find(id int64) (entity.Product, error) {
	db := util.DBConn()
	defer db.Close()

	var product = entity.Product{}

	query := "SELECT id, title, description FROM products WHERE id = ?;"
	row := db.QueryRow(query, id)
	err := row.Scan(&product.ID, &product.Title, &product.Description)
	if err != nil {
		return product, errors.New(fmt.Sprintf("error fetching product with ID %d from the database: %v", id, err))
	}

	return product, nil
}

func (*mysqlRepo) FindAll() (*[]entity.Product, error) {
	db := util.DBConn()
	defer db.Close()

	query := "SELECT id, title, description FROM products;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error fetching all products from the database: %v", err))
	}

	var products []entity.Product
	var product = entity.Product{}

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Title, &product.Description)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error fetching all products from the database: %v", err))
		}

		products = append(products, product)
	}

	return &products, nil
}

// SearchByTitle
func (*mysqlRepo) SearchByTitle(term string) (*[]entity.Product, error) {
	db := util.DBConn()
	defer db.Close()

	query := "SELECT id, title, description FROM products WHERE title LIKE ?;"
	rows, err := db.Query(query, "%"+term+"%")
	if err != nil {
		return nil, err
	}

	var products []entity.Product
	var product = entity.Product{}

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Title, &product.Description)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return &products, nil
}

func (*mysqlRepo) Save(product entity.Product) (*entity.Product, error) {
	db := util.DBConn()
	defer db.Close()

	query := "INSERT INTO products (title, description) VALUES(?, ?);"
	stmt, stmtErr := db.Prepare(query)
	if stmtErr != nil {
		return &entity.Product{}, errors.New(fmt.Sprintf("error saving product into the database: %v", stmtErr))
	}

	res, queryErr := stmt.Exec(product.Title, product.Description)
	if queryErr != nil {
		return &entity.Product{}, errors.New(fmt.Sprintf("error saving product into the database: %v", queryErr))
	}

	id, getLastInsertIDErr := res.LastInsertId()
	product.ID = id
	if getLastInsertIDErr != nil {
		return &entity.Product{}, errors.New(fmt.Sprintf("error saving product into the database: %v", getLastInsertIDErr))
	}

	return &product, nil
}
