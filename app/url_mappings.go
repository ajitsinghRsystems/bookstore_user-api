package app
import (
	"github.com/ajitsinghRsystems/bookstore_users-api/controllers/ping"
)
func mapUrls(){
	router.GET("/ping", ping.Ping)

}