package app

import (
	"github.com/nelbermora/bookstore_users-api/controllers/ping"
	"github.com/nelbermora/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/", users.SearchUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/internals/users/search", users.FindByStatus)

}
