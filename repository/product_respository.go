package repository

import (
	"database/sql"
	"go-products-api/model"

	"github.com/rs/zerolog/log"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "select id, name, price from products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		log.Error().Err(err).Stack().Msg("repository - error when select")
		return []model.Product{}, err
	}

	var products []model.Product
	var product model.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			log.Error().Err(err).Stack().Msg("repository - error when scan result")
			return []model.Product{}, err
		}
		products = append(products, product)
	}
	rows.Close()
	return products, nil
}
