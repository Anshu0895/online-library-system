package routes

import (
	"online-library-system/controllers"
	"online-library-system/middleware"

	"github.com/gin-gonic/gin"
)

func RequestRoutes(router *gin.Engine) {
	router.POST("/raise-request", middleware.RoleBasedAccessControl("Reader"), controllers.RaiseIssueRequest)
	router.GET("/requests", middleware.RoleBasedAccessControl("Admin"), controllers.GetRequestEvents)
	router.GET("/requests/:id", middleware.RoleBasedAccessControl("Admin"), controllers.GetRequestEventsByID)
	router.PUT("/requests/:id/approve", middleware.RoleBasedAccessControl("Admin"), controllers.ApproveIssueRequest)
	router.PUT("/requests/:id/reject", middleware.RoleBasedAccessControl("Admin"), controllers.RejectIssueRequest)
}
