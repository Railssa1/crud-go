package routes

import (
	"github.com/Railssa1/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

// Função responsável por inicializar rotas da aplicação
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/v1/user/id/:id", controller.FindUserById)
	r.GET("/v1/user/email/:email", controller.FindUserByEmail)
	r.POST("/v1/user", controller.CreateUser)
	r.PUT("/v1/user", controller.UpdateUser)
	r.DELETE("/v1/user/:id", controller.DeleteUser)
}
