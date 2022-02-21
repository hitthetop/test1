package dao

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var (
	DB *sqlx.DB
)
func InitDB() (err error){
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/users")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	DB = database
	err = DB.Ping()
	if err != nil{
		fmt.Printf("open %s faild, err:%v\n",err)
		return
	}
	fmt.Println("连接成功！")
	return
}