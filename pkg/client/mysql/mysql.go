package mysql

import (
	"database/sql"
)

func NewClient() *sql.DB {
	db, err := sql.Open("mysql", "sandbox_user:passpass@tcp(127.0.0.1:3306)/sandbox")
	if err != nil {
		panic(err)
	}
	return db
}
