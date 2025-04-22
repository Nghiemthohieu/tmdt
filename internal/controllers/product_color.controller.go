package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductColorController struct {
	ProductColorService *services.ProductColorService
}

func NewProductColorController() *ProductColorController {
	return &ProductColorController{
		ProductColorService: services.NewProductColorService(),
	}
}

func (c *ProductColorController) CreateProductColor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productColor models.ProductColor
		if err := ctx.ShouldBindJSON(&productColor); err != nil {
			response.ErrorRespone(ctx, 400, 40001, err)
			return
		}
		code, err := c.ProductColorService.CreateProductColor(&productColor)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product color created successfully")
	}
}

func (c *ProductColorController) GetProductColorByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			response.ErrorRespone(ctx, 400, 40001, nil)
			return
		}
		productColor, code, err := c.ProductColorService.GetProductColorByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, productColor)
	}
}

func (c *ProductColorController) GetAllProductColors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productColors, code, err := c.ProductColorService.GetAllProductColors()
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, productColors)
	}
}

func (c *ProductColorController) UpdateProductColor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productColor models.ProductColor
		if err := ctx.ShouldBindJSON(&productColor); err != nil {
			response.ErrorRespone(ctx, 400, 40001, err)
			return
		}
		code, err := c.ProductColorService.UpdateProductColor(&productColor)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product color updated successfully")
	}
}

func (c *ProductColorController) DeleteProductColor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			response.ErrorRespone(ctx, 400, 40001, nil)
			return
		}
		code, err := c.ProductColorService.DeleteProductColor(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product color deleted successfully")
	}
}
