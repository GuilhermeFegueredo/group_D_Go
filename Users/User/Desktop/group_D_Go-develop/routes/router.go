package routes

import (
	"api-produto/controllers"
	"api-produto/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.ProdutoServiceInterface) *gin.Engine {
	main := router.Group("api/v1")
	{
		produtos := main.Group("produtos")
		{
			produtos.GET("/:id", func(c *gin.Context) {
				controllers.GetProduto(c, service)
			})
			produtos.GET("/", func(c *gin.Context) {
				controllers.GetAll(c, service)
			})
			produtos.POST("/", func(c *gin.Context) {
				controllers.Create(c, service)
			})
			produtos.DELETE("/:id", func(c *gin.Context) {
				controllers.Delete(c, service)
			})
		}
	}

	return router
}