package main

import (
	"fmt"
	"log"

	"github.com/Railssa1/crud-go/routes"
	"github.com/Railssa1/crud-go/src/config/database/mongodb"
	"github.com/Railssa1/crud-go/src/controller"
	"github.com/Railssa1/crud-go/src/repository"
	"github.com/Railssa1/crud-go/src/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	app := gin.Default()

	// Carregando váriaveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Conectando ao banco de dados
	database, err := mongodb.InitConnection()
	if err != nil {
		log.Fatalf("Erro trying to connect to databse, error=%s \n", err.Error())
	}

	// Iniciando o repositorio
	userRepository := repository.NewUserRepository(database)

	// Iniciando service
	service := service.NewUserDomainService(userRepository)

	// Iniciando controller
	userController := controller.NewUserController(service)

	routes.InitRoutes(&app.RouterGroup, userController)

	if err := app.Run(":8080"); err != nil {
		fmt.Println("Erro ao iniciar aplicação", err)
	}
}
