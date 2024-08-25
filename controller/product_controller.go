package controller

import (
	"go-products-api/model"
	"go-products-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		log.Error().Err(err).Stack().Msg("products controller - error when searching")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdProdutct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdProdutct)

}

func (p *productController) GetProductsById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	log.Debug().Msg("looking for product with id " + paramId)
	if paramId == "" {
		response := model.Response{
			Message: "id not be empty or null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(paramId)
	if err != nil {
		response := model.Response{
			Message: "id need to be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := p.productUsecase.GetProductById(id)

	if err != nil {
		log.Error().Err(err).Stack().Msg("products controller - error when searching by id")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "No Products were found with this identifier",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
