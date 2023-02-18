package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang?collation=utf8mb4_unicode_ci")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(150)
	db.SetMaxOpenConns(1500)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
