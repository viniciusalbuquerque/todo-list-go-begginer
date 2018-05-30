package mydb
import "database/sql"

var db *sql.DB

const (
	dbName = "todoServerDB"
	conn = ""
)

func OpenConnection() {
	dbT, err := sql.Open(conn, dbName)
        if err != nil {
		panic(err)
	}

	db = dbT
}
