package main

import (
	"fmt"

	"github.com/Railssa1/crud-go/routes"
	"github.com/Railssa1/crud-go/src/config/database/mongodb"
	"github.com/Railssa1/crud-go/src/controller"
	"github.com/Railssa1/crud-go/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// Conectando ao banco de dados
	mongodb.InitConnection()

	// Iniciando service
	service := service.NewUserDomainService()

	// Iniciando controller
	userController := controller.NewUserController(service)

	routes.InitRoutes(&app.RouterGroup, userController)

	if err := app.Run(":8080"); err != nil {
		fmt.Println("Erro ao iniciar aplicação", err)
	}
}
