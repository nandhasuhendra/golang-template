package config

import (
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
	"os"
	"golang-template/helper"
	"time"
)


func NewDB() *sql.DB {
  LoadEnv()

  db, err := sql.Open("mysql", os.Getenv("MYSQL_URL") + "?multiStatements=true")
  helper.PanicIfError(err)

  db.SetMaxIdleConns(5)
  db.SetMaxOpenConns(20)
  db.SetConnMaxLifetime(60 * time.Minute)
  db.SetConnMaxIdleTime(10 * time.Minute)

  return db
}
