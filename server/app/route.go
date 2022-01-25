package app

/*
middleware enable us to make API calls
 between  server and  client(frontend)
*/
import (
	"username-finder/server/controller"
	"username-finder/server/middleware"
)

func route() {
	router.Use(middleware.CORSMiddleware()) //to enable api request between client and server

	router.POST("/username", controller.Username)
	router.POST("/qr", controller.QRcodegenerator)
}
