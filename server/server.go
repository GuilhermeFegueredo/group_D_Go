package server

import (
	"api-produto/config"
	"api-produto/routes"
	"api-produto/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	SRV_PORT string
	SERVER   *gin.Engine
}

func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

func (s *Server) Run(service service.ProdutoServiceInterface) {
	router := routes.ConfigRoutes(s.SERVER, service)

	log.Print("Server is running at port: ", s.SRV_PORT)
	log.Fatal(router.Run(":" + s.SRV_PORT))
}
