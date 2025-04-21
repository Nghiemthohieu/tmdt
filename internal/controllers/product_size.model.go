package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductSizeController struct {
	ProductSizeService *services.ProductSizeService
}

func NewProductSizeController(productSizeService *services.ProductSizeService) *ProductSizeController {
	return &ProductSizeController{
		ProductSizeService: productSizeService,
	}
}

func (c *ProductSizeController) CreateProductSize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productSize models.ProductSize
		if err := ctx.ShouldBindJSON(&productSize); err != nil {
			response.ErrorRespone(ctx, 400, 40001, err)
			return
		}
		code, err := c.ProductSizeService.CreateProductSize(&productSize)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product size created successfully")
	}
}

func (c *ProductSizeController) GetProductSizeByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			response.ErrorRespone(ctx, 400, 40001, nil)
			return
		}
		productSize, code, err := c.ProductSizeService.GetProductSizeByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, productSize)
	}
}

func (c *ProductSizeController) GetAllProductSizes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productSizes, code, err := c.ProductSizeService.GetAllProductSizes()
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, productSizes)
	}
}

func (c *ProductSizeController) UpdateProductSize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productSize models.ProductSize
		if err := ctx.ShouldBindJSON(&productSize); err != nil {
			response.ErrorRespone(ctx, 400, 40001, err)
			return
		}
		code, err := c.ProductSizeService.UpdateProductSize(&productSize)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product size updated successfully")
	}
}

func (c *ProductSizeController) DeleteProductSize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			response.ErrorRespone(ctx, 400, 40001, nil)
			return
		}
		code, err := c.ProductSizeService.DeleteProductSize(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product size deleted successfully")
	}
}
