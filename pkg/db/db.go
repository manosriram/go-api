package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var Connection *sql.DB;
func Initialize() *sql.DB {
    db, err := sql.Open("mysql", "root:password@/gouser");
    Connection = db;
    if err != nil {
	log.Fatal(err);
    } 
    return Connection;
}

func GetConnection() *sql.DB {
    return Connection;
}
