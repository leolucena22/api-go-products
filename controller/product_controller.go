package controller

import (
	usecase "api/Usecase"
	"api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productsUsecase usecase.ProductsUsecase
}

func NewProductController(usecase usecase.ProductsUsecase) ProductController {
	return ProductController{
		productsUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productsUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productsUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	product, err := p.productsUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)

		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) UpdatePriceProduct(ctx *gin.Context) {
	// Capturar ID
	id := ctx.Param("productId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID é obrigatório"})
		return
	}

	// Converter ID
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID inválido"})
		return
	}

	// Capturar novo preço
	var request struct {
		Price float64 `json:"price"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Dados inválidos"})
		return
	}

	// Atualizar
	product, err := p.productsUsecase.UpdatePriceProduct(productId, request.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno"})
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Produto não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
