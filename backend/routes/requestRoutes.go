package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
)

func RequestRoutes(router *gin.Engine) {
	router.POST("/requests", controllers.CreateRequestEvent)
	router.GET("/requests", controllers.GetRequestEvents)
	router.PUT("/requests/:id/approve", controllers.ApproveRequestEvent)
}
