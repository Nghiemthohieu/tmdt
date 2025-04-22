package user

import (
	"github.com/gin-gonic/gin"
)

type ProductSizeRouter struct {
}

func (p *ProductSizeRouter) InitProductSizeRouter(Router *gin.RouterGroup) {
	productRouterPublic := Router.Group("/product_size")
	{
		productRouterPublic.GET("/list")
	}
}
