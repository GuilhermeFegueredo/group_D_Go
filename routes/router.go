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
			produtos.PUT("/:id", func(c *gin.Context) {
				controllers.Update(c, service)
			})
		}
	}

	return router
}
