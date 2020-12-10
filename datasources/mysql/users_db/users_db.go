package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_db_user   = "root"
	mysql_db_host   = "localhost:3306"
	mysql_db_schema = "envioskyl"
)

// seteo la variable global, la db se abre solo una vez
var (
	UsersDB *sql.DB
	/*
		username = os.Getenv(mysql_db_user)
		host     = os.Getenv(mysql_db_host)
		schema   = os.Getenv(mysql_db_schema)
	*/
)

func init() {
	dataSourceName := fmt.Sprintf("%s@tcp(%s)/%s", mysql_db_user, mysql_db_host, mysql_db_schema)
	var err error
	UsersDB, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}
	err = UsersDB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database Configured")
}
