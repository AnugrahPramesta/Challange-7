package route

import (
	"challange-8/handler"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	api := r.Group("/books") // prefix
	{
		api.GET("", server.GetBooks)          // /employees
		api.POST("", server.Addbook)          // /employees
		api.GET("/:id", server.GetBookByID)   // /employee/:id
		api.PUT("/:id", server.UpdateBook)    // /employee/:id
		api.DELETE("/:id", server.DeleteBook) // /employee/:id
	}
}
