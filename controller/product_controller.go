package controller

import (
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUseCase: usecase,
	}
}

// Recupera Produtos
func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.ProductUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	// Retorna o json com os produtos, resposta 200
	ctx.JSON(http.StatusOK, products)
}
