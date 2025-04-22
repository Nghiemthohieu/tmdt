package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

func (p *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouterPublic := Router.Group("/category")
	{
		CategoryRouterPublic.GET("/list", controllers.NewCategoryController().GetAllCategories())
		CategoryRouterPublic.GET("/get/:id", controllers.NewCategoryController().GetCategoryByID())
	}
}
