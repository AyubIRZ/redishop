package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // to be used with native sql driver
	"os"
)

// DBConn opens and returns a new connection to the database(from a connection pool)
func DBConn() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":3306)/" + os.Getenv("MYSQL_DATABASE") + "?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}

