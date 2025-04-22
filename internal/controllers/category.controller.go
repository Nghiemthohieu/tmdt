package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		CategoryService: services.NewCategoryService(),
	}
}

func (c *CategoryController) CreateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category models.Category
		if err := ctx.ShouldBindJSON(&category); err != nil {
			response.ErrorRespone(ctx, 400, 40001, err)
			return
		}
		code, err := c.CategoryService.CreateCategory(&category)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Category created successfully")
	}
}

func (c *CategoryController) GetCategoryByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			response.ErrorRespone(ctx, 400, 40001, nil)
			return
		}
		category, code, err := c.CategoryService.GetCategoryByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, category)
	}
}

func (c *CategoryController) GetAllCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories, code, err := c.CategoryService.GetAllCategories()
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, categories)
	}
}

func (c *CategoryController) UpdateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category models.Category
		if err := ctx.ShouldBindJSON(&category); err != nil {
			response.ErrorRespone(ctx, 400, 40001, err)
			return
		}
		code, err := c.CategoryService.UpdateCategory(&category)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Category updated successfully")
	}
}

func (c *CategoryController) DeleteCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			response.ErrorRespone(ctx, 400, 40001, nil)
			return
		}
		code, err := c.CategoryService.DeleteCategory(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Category deleted successfully")
	}
}
