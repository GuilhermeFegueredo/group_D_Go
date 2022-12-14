package main

import (
	"api-produto/config"
	"api-produto/database"
	"api-produto/routes"
	"api-produto/server"
	"api-produto/service"
	"api-produto/webui"
	"encoding/json"
	"os"
)

func main() {

	default_conf := &config.Config{}

	if file_config := "config.json"; file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	conf := config.NewConfig(default_conf)

	dbpool := database.NewDB(conf)

	service := service.NewProdutoService(dbpool)
	serv := server.NewServer(conf)

	// Inicia as rotas
	router := routes.ConfigRoutes(serv.SERVER, service)

	if conf.WEB_UI {
		webui.RegisterUIHandlers(router)
	}

	// Registra as rotas
	server.Run(router, serv, service)

}
