package app

import (
	"example.com/hello/github.com/miftahulrespati/bookstore_users-api/controllers/ping"
	"example.com/hello/github.com/miftahulrespati/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping) // function not called but defined

	router.GET("/users/:user_id", users.GetUser) // must add (c *gin.Context) in controllers
	router.POST("/users", users.CreateUser)
}
