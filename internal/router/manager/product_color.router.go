package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductColorRouter struct{}

func (p *ProductColorRouter) InitProductColorRouter(Router *gin.RouterGroup) {
	ProductColorRouterPrivate := Router.Group("/admin/product_color")
	{
		ProductColorRouterPrivate.GET("/list", controllers.NewProductColorController().GetAllProductColors())
		ProductColorRouterPrivate.POST("/add", controllers.NewProductColorController().CreateProductColor())
		ProductColorRouterPrivate.PUT("/update", controllers.NewProductColorController().UpdateProductColor())
		ProductColorRouterPrivate.DELETE("/delete/:id", controllers.NewProductColorController().DeleteProductColor())
		ProductColorRouterPrivate.GET("/get/:id", controllers.NewProductColorController().GetProductColorByID())
	}
}
