package main

import (
	"github.com/disaster_management_backend/config"
	"github.com/disaster_management_backend/internals/database"
	"github.com/disaster_management_backend/internals/routes"
	"fmt"
)

func main() {
	config.LoadConfig()

	fmt.Println("env values :" , config.AppConfig.DatabaseURL)

	_ = database.InitDB(config.AppConfig.DatabaseURL)
	r := routes.SetupRouter()
	fmt.Printf("Server started on port %s\n", config.AppConfig.Port)
	r.Run(fmt.Sprintf(":%s", config.AppConfig.Port))
}
