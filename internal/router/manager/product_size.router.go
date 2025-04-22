package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductSizeRouter struct{}

func (p *ProductSizeRouter) InitProductSizeRouter(Router *gin.RouterGroup) {
	ProductSizeRouterPrivate := Router.Group("/admin/product_size")
	{
		ProductSizeRouterPrivate.GET("/list", controllers.NewProductSizeController().GetAllProductSizes())
		ProductSizeRouterPrivate.POST("/add", controllers.NewProductSizeController().CreateProductSize())
		ProductSizeRouterPrivate.PUT("/update", controllers.NewProductSizeController().UpdateProductSize())
		ProductSizeRouterPrivate.DELETE("/delete/:id", controllers.NewProductSizeController().DeleteProductSize())
		ProductSizeRouterPrivate.GET("/get/:id", controllers.NewProductSizeController().GetProductSizeByID())
	}
}
