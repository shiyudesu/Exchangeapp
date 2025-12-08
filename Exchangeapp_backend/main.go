package main

import (
	"exchangeapp/Exchangeapp/Exchangeapp_backend/config"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Println(config.AppConfig.App.Port)
}
