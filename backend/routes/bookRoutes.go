package routes

import (
	"online-library-system/controllers"
	"online-library-system/middleware"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	router.POST("/books", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.AddBook)
	router.PUT("/books/:isbn", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.UpdateBook)
	router.DELETE("/books/:isbn", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.RemoveBook)
	router.GET("/books", middleware.RoleBasedAccessControl("Reader"), controllers.GetBooks)
	router.GET("/book", middleware.RoleBasedAccessControl("LibraryAdmin"), controllers.GetBooks)
	router.GET("/books/:isbn", middleware.RoleBasedAccessControl("Reader"), controllers.GetBook)
	router.GET("/books/search", middleware.RoleBasedAccessControl("Reader"), controllers.SearchBooks)
}
