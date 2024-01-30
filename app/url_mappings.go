package app

import (
	"github.com/vijaykatta6/bookstore-users-api/controllers/ping"
	"github.com/vijaykatta6/bookstore-users-api/controllers/users"
)

func mapUrls() {
	// for this endpoint this particular controller should call
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
