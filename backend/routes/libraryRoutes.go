package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
)

func LibraryRoutes(router *gin.Engine) {
	router.POST("/libraries", controllers.CreateLibrary)
}
