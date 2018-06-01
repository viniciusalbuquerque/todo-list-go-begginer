package mydb
import (
		"database/sql"
		"fmt"
		_ "github.com/lib/pq"
		)

var DB *sql.DB

const (
	conn = "postgres"
	DB_USER     = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME     = "todoserverdb"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func OpenConnection() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
	DB, err = sql.Open(conn, dbinfo)
	checkErr(err)
    // defer DB.Close()

	err = DB.Ping(); 
	checkErr(err)

}
