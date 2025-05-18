package main

import (
	"fmt"

	"github.com/Railssa1/crud-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.InitRoutes(&app.RouterGroup)

	if err := app.Run(":8080"); err != nil {
		fmt.Println("Erro ao iniciar aplicação", err)
	}
}
