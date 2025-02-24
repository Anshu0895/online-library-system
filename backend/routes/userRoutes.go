package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateAdmin)
	router.GET("/users", controllers.GetUsers)
	router.GET("/admins", controllers.GetAdmins)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
