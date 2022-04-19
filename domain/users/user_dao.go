package users

import (
	"bookstore_user-api/datasources/users_db"
	"bookstore_user-api/utils/date_utils"
	"bookstore_user-api/utils/errors"
	"fmt"
	"log"
	"context"
	
)

var (
	usersDB =make(map[int64] *User)
)
const(
	queryInserUser="INSERT INTO users(first_name, last_name, email, date_created) VALUES(@p1, @p2, @p3, @p4);select isNull(SCOPE_IDENTITY(), -1)"
	queryGetUser="Select ID, First_name,Last_name, Email, Date_Created from Users where ID =@p1;"
	queryUpdateUser="Update users set first_name=@p1 , last_name=@p2, email=@p3 from Users where ID =@p4;"
)
func (user *User) Get()(*errors.RestErr){

	if err := users_db.Client.Ping(); err != nil{
		panic(err)
	}
	
	stmt,err :=users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

 	result := stmt.QueryRow(user.Id)
 	if  err:= result.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated); err != nil{
		 fmt.Print()
		 return errors.NewNotFoundError(fmt.Sprintf("User %d not found %s", user.Id,err.Error()))
 	}

	// user.Id=result.Id
	 //user.FirstName=result.FirstName
	 //user.LastName=result.LastName
	 //user.Email=result.Email
	 //user.DateCreated= result.DateCreated
	return nil
}
func  (user *User) Save() *errors.RestErr {
	ctx := context.Background()

	stmt,err :=users_db.Client.Prepare(queryInserUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	//

	user.DateCreated = date_utils.GetNowString()
	insertResult, err :=stmt.QueryContext(ctx,user.FirstName, user.LastName, user.Email, user.DateCreated)
	//insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		//logger.Error("error when trying to save user", saveErr)
		log.Println(err)
		
		return errors.NewInternalServerError(
			fmt.Sprintf("error when tying to save user %s", err.Error()))
	}
		var newID int64
		
		insertResult.Next() 
		err = insertResult.Scan(&newID)
			if err != nil {
				//logger.Error("error when trying to get last insert id after creating a new user", err)
				return errors.NewInternalServerError(fmt.Sprintf("error when tying to save user%s", err.Error()))
			}
		
	
	user.Id = newID
	usersDB[user.Id] = user
	return nil
	//current := usersDB[user.Id]
	//if current != nil{
	//	
	//	if current.Email == user.Email{
	///		return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered",user.Email))
		//}

		//return errors.NewBadRequestError(fmt.Sprintf("User %d already exist",user.Id))

	//}
	
	//usersDB[user.Id] = user
	
	//return nil
}
func (user *User) Update() *errors.RestErr{
stmt, err := users_db.Client.Prepare((queryUpdateUser))
if err!= nil{
	return errors.NewInternalServerError(err.Error())

}
defer stmt.Close()
result , err := stmt.Exec(user.FirstName,user.LastName, user.Email, user.Id)

if err !=nil{
	return errors.NewInternalServerError(err.Error())
}
log.Println(result.RowsAffected())
return nil
	
}