package routes

import (
	"github.com/gin-gonic/gin"
	"online-library-system/controllers"
)

func BookRoutes(router *gin.Engine) {
	router.POST("/books", controllers.AddBook)
	router.PUT("/books/:isbn", controllers.UpdateBook)
	router.DELETE("/books/:isbn", controllers.RemoveBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:isbn", controllers.GetBook)
}
