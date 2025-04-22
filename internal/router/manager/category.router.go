package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (p *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouterPrivate := Router.Group("/admin/category")
	{
		CategoryRouterPrivate.GET("/list", controllers.NewCategoryController().GetAllCategories())
		CategoryRouterPrivate.POST("/add", controllers.NewCategoryController().CreateCategory())
		CategoryRouterPrivate.PUT("/update", controllers.NewCategoryController().UpdateCategory())
		CategoryRouterPrivate.DELETE("/delete/:id", controllers.NewCategoryController().DeleteCategory())
		CategoryRouterPrivate.GET("/get/:id", controllers.NewCategoryController().GetCategoryByID())
	}
}
