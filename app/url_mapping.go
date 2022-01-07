package app

import (
	"github.com/wandyirawan/bookstore_users-api/controllers/ping"
	"github.com/wandyirawan/bookstore_users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/ysers/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
