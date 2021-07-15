package app

import (
	"github.com/pavangujar23/bookStore_UsersAPI/controllers/ping"
	"github.com/pavangujar23/bookStore_UsersAPI/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)

	// router.GET("/users", controllers.SearchUser)
	router.POST("/users", users.CreateUser)

}
