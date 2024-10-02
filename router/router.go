package router

import (
	"github.com/gin-gonic/gin"
)

//
func Initializer() {
	
	// Initializer Router
	router := gin.Default()
	
	initializeRoutes(router)

	// Run server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
	
