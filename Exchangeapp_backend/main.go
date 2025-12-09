package main

import (
	"exchangeapp/Exchangeapp/Exchangeapp_backend/config"
	"exchangeapp/Exchangeapp/Exchangeapp_backend/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(config.AppConfig.App.Port) // listen and serve on 0.0.0.0:8080
}
