package usecase

import (
	"api/model"
	"api/repository"

	_ "github.com/lib/pq"
)

type ProductsUsecase struct {
	repository repository.ProductsRepository
}

func NewProductUsecase(repo repository.ProductsRepository) ProductsUsecase {
	return ProductsUsecase{
		repository: repo,
	}
}

func (pu *ProductsUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductsUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductsUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductsUsecase) UpdatePriceProduct(id_product int, newPrice float64) (*model.Product, error) {
	return pu.repository.UpdatePriceProduct(id_product, newPrice)
}
