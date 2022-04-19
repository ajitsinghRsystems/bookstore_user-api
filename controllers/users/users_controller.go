package users

import (
	"bookstore_user-api/domain/users"
	"bookstore_user-api/services"
	"bookstore_user-api/utils/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
var(
	counter int
)
func getUserId(userIDParam string) (int64, *errors.RestErr){
	userID,userErr:=strconv.ParseInt(userIDParam,10,64)
	if userErr !=nil{
		return 0, errors.NewBadRequestError("User id should be number")
	
	}
	return userID,nil
}
func GetUser(c *gin.Context){
	userID,userErr:= getUserId(c.Param("user_id"))
	if userErr !=nil{
		err:= errors.NewBadRequestError("User id should be number")
		c.JSON(err.Status,err)
		return
	}
	user,getErr :=services.GetUser(userID)
	if getErr !=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusCreated, user)

}
func CreateUsers( c *gin.Context){
	var user users.User
	fmt.Println(user)
	if err:=c.ShouldBindJSON(&user); err != nil{
		
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status,restErr)

		return
	}
	result,saveErr :=services.CreateUser(user)
	if saveErr !=nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}
	
	c.JSON(http.StatusCreated, result)
}
func UpdateUser(c *gin.Context){

	userID,userErr:=strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr !=nil{
		err:= errors.NewBadRequestError("User id should be number")
		c.JSON(err.Status,err)
		return
	}

	var user users.User
	fmt.Println(user)
	if err:=c.ShouldBindJSON(&user); err != nil{
		
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status,restErr)

		return
	}	
	user.Id=userID
	isPartial := c.Request.Method== http.MethodPatch

	result,saveErr :=services.UpdateUser(isPartial,user)
	if saveErr !=nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}
	
	c.JSON(http.StatusOK, result)
}
func DeleteUser(c *gin.Context){
	userID,userErr:= getUserId(c.Param("user_id"))
	if userErr !=nil{
		err:= errors.NewBadRequestError("User id should be number")
		c.JSON(err.Status,err)
		return
	}
	if err:=services.DeleteUser(userID); err != nil{
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK, map[string] string{"status":"deleted"})
}