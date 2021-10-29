package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default() // router should only exist in the app package
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
