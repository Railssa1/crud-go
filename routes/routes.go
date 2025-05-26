package routes

import (
	"github.com/Railssa1/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

// Função responsável por inicializar rotas da aplicação
func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/v1/user/id/:id", userController.FindUserById)
	r.GET("/v1/user/email/:email", userController.FindUserByEmail)
	r.POST("/v1/user", userController.CreateUser)
	r.PUT("/v1/user", userController.UpdateUser)
	r.DELETE("/v1/user/:id", userController.DeleteUser)
}
