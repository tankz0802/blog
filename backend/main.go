package main

import (
	"blog/config"
	"blog/model"
	"blog/router"
)

func main() {
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter()
}
