package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marciopriebe/crud-go.git/handler"
)

func initializeRoutes(router *gin.Engine)  {
	v1 := router.Group("/api/v1")
	{
		// Create 
		v1.GET("/opening", handler.CreateOpeningHandler)

		// Show
		v1.POST("/opening", handler.ShowOpeningHandler)

		// DELETE
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		
		//UPDATE
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// SHOW ALL 
		v1.GET("/openings", handler.ListOpeningsHandler)

	}
}