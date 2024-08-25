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

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	query, err := pr.connection.Prepare(`
		insert into products (name, price) values ($1, $2) returning id
	`)
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to prepare product creation")
		return product, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to associate parameters to product creation query")
		return product, err
	}
	query.Close()

	return product, err
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare(`
		select id, name, price from products where id = $1
	`)
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to prepare query to find product by id")
		return nil, err
	}
	var product model.Product
	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to execute query to find product by id")
		if err == sql.ErrNoRows {
			log.Debug().Msg("not found product")
			return nil, nil
		}
		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *ProductRepository) DeleteProductById(id int) (*model.Product, error) {
	product, err := pr.GetProductById(id)
	if err != nil {
		return nil, err
	}
	query, err := pr.connection.Prepare(`
		delete from products where id = $1
	`)
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to prepare query to delete product by id")
		return nil, err
	}

	err = query.QueryRow(id).Err()
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to execute query to delete product by id")
		return nil, err
	}
	return product, nil

}
