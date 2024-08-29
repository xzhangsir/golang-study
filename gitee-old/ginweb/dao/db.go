package dao;

import (
	// "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

func InitDB()(err error){
	// 数据库信息
	dsn := "root:haosql@tcp(127.0.0.1:3306)/ginweb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	return 
}
func Close() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}