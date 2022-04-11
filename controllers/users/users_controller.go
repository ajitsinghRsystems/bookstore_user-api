package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
var(
	counter int
)
func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Implement Me")
}
func CreateUsers( c *gin.Context){
	c.String(http.StatusNotImplemented, "Implement Me")
}
