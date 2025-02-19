package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

}
