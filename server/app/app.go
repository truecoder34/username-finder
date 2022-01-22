package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {

	route()

	port := os.Getenv("PORT") //using heroku host
	if port == "" {
		port = "8888" //localhost
	}

	router.Run(":" + port)
}
