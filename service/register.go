package service

import (
	"encoding/json"
	"fmt"
	"myproject/dao"
	"myproject/model"
)

func Register(name string,password string,emil string,introduction string,phone int,qq int,
	gender string,birth string) string {
	var u model.User
	sqlStr := "select name from user where name = ?"
	rowOBJ:= dao.DB.QueryRow(sqlStr,name)
	rowOBJ.Scan(&u.UserName)
	if name==u.UserName{
		fmt.Println("账号重复")
		return ""
	}
	sqlstr:= `insert into user(name,password,introduction,emil,phone,qq,gender,birth)
    values(?,?,?,?,?,?,?,?)`
	_,err :=dao.DB.Exec(sqlstr,name,password,introduction,emil,phone,qq,gender,birth)
	if err!=nil {
		fmt.Println(err)
	}
	//建立一张表去保存用户的收藏和关注信息
	str := "create table %s (name varchar(255),topicName varchar(255),dislikeName varchar(255))"
	sql :=fmt.Sprintf(str,freshtoken.UserName)
	_,err =dao.DB.Exec(sql)
	if err!=nil {
		fmt.Println(err)
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}