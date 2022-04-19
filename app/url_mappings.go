package app

import "bookstore_user-api/controllers/ping"
import "bookstore_user-api/controllers/users"
func mapUrls(){
	router.GET("/ping",ping.Ping)

	router.GET("/users/:user_id",users.GetUser)
	router.POST("/users", users.CreateUsers)
	router.PUT("/users/:user_id",users.UpdateUser)
	router.PATCH("/users/:user_id",users.UpdateUser)
}
