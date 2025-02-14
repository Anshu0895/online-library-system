package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
	"online-library-system/middleware"
)

func IssueRoutes(router *gin.Engine) {
	router.POST("/issues", middleware.RoleBasedAccessControl("Admin"), controllers.CreateIssueRegistry)
	router.GET("/issues", middleware.RoleBasedAccessControl("Admin"), controllers.GetIssueRegistries)
	router.PUT("/issues/:id", middleware.RoleBasedAccessControl("Admin"), controllers.UpdateIssueStatus)
}
