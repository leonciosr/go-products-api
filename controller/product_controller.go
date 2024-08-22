package controller

import (
	"go-products-api/usecase"
	"net/http"

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
	}

	ctx.JSON(http.StatusOK, products)
}
