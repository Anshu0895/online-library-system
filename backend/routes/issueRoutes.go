import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
)

func IssueRoutes(router *gin.Engine) {
	router.POST("/issues", controllers.CreateIssueRegistry)
	router.GET("/issues", controllers.GetIssueRegistries)
	router.PUT("/issues/:id", controllers.UpdateIssueStatus)
}