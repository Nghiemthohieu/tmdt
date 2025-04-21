package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	// r.Use()

	v1 := r.Group("/api")
	{
		v1.GET("/ping")
	}
	return r
}
