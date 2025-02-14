package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
	"online-library-system/middleware"
)

func RequestRoutes(router *gin.Engine) {
	router.POST("/requests", middleware.RoleBasedAccessControl("Reader"), controllers.CreateRequestEvent)
	router.GET("/requests", middleware.RoleBasedAccessControl("Admin"), controllers.GetRequestEvents)
	router.PUT("/requests/:id/approve", controllers.ApproveRequestEvent)
}
