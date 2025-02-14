package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
	"online-library-system/middleware"
)

func BookRoutes(router *gin.Engine) {
	router.POST("/books", middleware.RoleBasedAccessControl("Admin"), controllers.AddBook)
	router.PUT("/books/:isbn", middleware.RoleBasedAccessControl("Admin"), controllers.UpdateBook)
	router.DELETE("/books/:isbn", middleware.RoleBasedAccessControl("Admin"), controllers.RemoveBook)
	router.GET("/books", middleware.RoleBasedAccessControl("Reader"), controllers.GetBooks)
	router.GET("/books/:isbn", middleware.RoleBasedAccessControl("Reader"), controllers.GetBook)
	router.GET("/books/search", middleware.RoleBasedAccessControl("Reader"), controllers.SearchBooks)
}
