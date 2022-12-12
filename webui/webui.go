package webui

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Funciona o front, mas não traz joga as informações no front

func RegisterUIHandlers(router *gin.Engine) {

	router.LoadHTMLGlob("./webui/dist/spa/*.html")

	router.Use(static.Serve("/webui", static.LocalFile("./webui/dist/spa/", true)))
	router.Use(static.Serve("/webui/assets", static.LocalFile("./webui/dist/spa/assets/", true)))
	router.Use(static.Serve("/webui/icons", static.LocalFile("./webui/dist/spa/icons/", true)))

	router.GET("/webui", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
