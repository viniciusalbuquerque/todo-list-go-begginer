package mydb
import "database/sql"

var db *sql.DB

const (
	dbName = "todoServerDB"
	conn = "postgres"
)

func OpenConnection() {
	var err error
	db, err = sql.Open(conn, dbName)
    if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
        panic(err)
    }
}
