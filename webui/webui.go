package webui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var server embed.FS

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func RegisterUIHandlers(router *gin.Engine) {

	router.LoadHTMLGlob("./webui/dist/spa/*.html")
	router.Use(static.Serve("/webui", EmbedFolder(server, "webui")))

	router.GET("/webui/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
