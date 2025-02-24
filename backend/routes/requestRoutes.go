package routes

import (
	"online-library-system/controllers"
	"online-library-system/middleware"

	"github.com/gin-gonic/gin"
)

func RequestRoutes(router *gin.Engine) {
	router.POST("/raise-request", middleware.RoleBasedAccessControl("Reader"), controllers.RaiseIssueRequest)
	router.GET("/requests", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.GetRequestEvents)
	router.GET("/pending-requests", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.GetPendingRequests)
	router.GET("/requests/:id", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.GetRequestEventsByID)
	router.PUT("/requests/:id/approve", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.ApproveIssueRequest)
	router.PUT("/requests/:id/reject", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.RejectIssueRequest)
}
