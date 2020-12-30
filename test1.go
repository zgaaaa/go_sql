package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// db全局变量
var db *sql.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/myemployees"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("连接失败 err = %v", err)
	}

}
