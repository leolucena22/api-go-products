package repository

import (
	"api/model"
	"database/sql"
	"fmt"
)

type ProductsRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductsRepository {
	return ProductsRepository{
		connection: connection,
	}
}

func (pr *ProductsRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductsRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductsRepository) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

func (pr *ProductsRepository) UpdatePriceProduct(id_product int, newPrice float64) (*model.Product, error) {
	query, err := pr.connection.Prepare(`
			UPDATE product 
			SET price = $1 
			WHERE id = $2 
			RETURNING id, product_name, price
	`)

	if err != nil {
		return nil, fmt.Errorf("erro ao preparar query: %w", err)
	}
	defer query.Close()

	var product model.Product
	err = query.QueryRow(newPrice, id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao atualizar produto: %w", err)
	}

	return &product, nil
}
