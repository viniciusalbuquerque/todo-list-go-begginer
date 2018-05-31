package mydb
import "database/sql"

var DB *sql.DB

const (
	dbName = "todoServerDB"
	conn = "postgres"
)

func OpenConnection() {
	var err error
	DB, err = sql.Open(conn, dbName)
    if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
        panic(err)
    }
}
