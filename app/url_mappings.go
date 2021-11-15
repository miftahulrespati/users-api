package app

import (
	"example.com/hello/github.com/miftahulrespati/users-api/controllers/ping"
	"example.com/hello/github.com/miftahulrespati/users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping) // function not called but defined

	// must add (c *gin.Context) in controllers
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
