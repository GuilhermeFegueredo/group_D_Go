package main

import (
	"api-produto/config"
	"api-produto/database"
	"api-produto/server"
	"api-produto/service"
	"encoding/json"
	"log"
	"os"
)

func main() {

	default_conf := &config.Config{}

	if file_config := "db_teste.json"; file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	conf := config.NewConfig(default_conf)
	dbpool := database.NewDB(conf)
	if dbpool != nil {
		log.Print("Se conectou ", dbpool.GetDB())
	}

	service := service.NewProdutoService(dbpool)

	server := server.NewServer(conf)
	server.Run(service)
}
