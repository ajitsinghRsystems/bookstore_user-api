package users_db

import (
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"fmt"
	"log"
	"context"
)
var(
	Client *sql.DB
	server = "LAPTOP-R9SOT732"
	port = 64206
	user = "sa"
 	password = "Password1"
 	database = "EmployeeDB"
	QueryContext context.Context
)

func init() {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",server, user, password, port, database)
	Client, err = sql.Open("sqlserver",connString)
	if err !=nil{
		panic(err)
	}
	if err = Client.Ping(); err != nil{
		panic(err)
	}

	log.Println("Database successfully connected")
}