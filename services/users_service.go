package services

import (
	"bookstore_user-api/domain/users"
	"bookstore_user-api/utils/errors"
	"bookstore_user-api/utils/crypto_utils"
	
)

func GetUser(userId int64) (*users.User,*errors.RestErr) {
	result:= &users.User{Id:userId}
	if err:=result.Get();err !=nil{
		return nil,err
	}
	return result,nil

}
func CreateUser(user users.User) (*users.User,*errors.RestErr){
if err:= user.Validate(); err !=nil{
	return nil,err
}
user.Password= crypto_utils.GetMd5(user.Password)
if err:= user.Save(); err !=nil{
return nil, err
}
	return &user,nil	
}
func UpdateUser(isPartial bool, user users.User) (*users.User,*errors.RestErr){
	current,err := GetUser(user.Id)
	if  err!=nil{
		return nil,err
	}
	if err := user.Validate() ; err!=nil{
		return nil,err
	}
	
	if !isPartial{
	current.FirstName=user.FirstName
	current.LastName=user.LastName
	current.Email=user.Email
	}else{

		if user.FirstName !=""{
			current.FirstName=user.FirstName
		}
		if user.LastName !=""{
			current.LastName=user.LastName
		}
		if user.Email !=""{
			current.Email=user.Email
		}
	}

	if err:= current.Update(); err!=nil {
		return nil, err
	}
	return current, nil

}
func DeleteUser(UserId int64)  (*errors.RestErr){
	user:= &users.User{Id:UserId}
	return user.Delete()

}
func Search(status string) ([]users.User,*errors.RestErr){
dao:= &users.User{}
return dao.FindByStatus(status)

}