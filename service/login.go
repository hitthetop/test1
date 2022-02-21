package service

import (
	"encoding/json"
	"fmt"
	"myproject/dao"
	"myproject/model"
)

var freshtoken model.FreshToken
func Login(name string,password string) string {
	var u model.User
	sqlStr := "select name,password from user where name = ?"
	rowOBJ:= dao.DB.QueryRow(sqlStr,name)
	err := rowOBJ.Scan(&u.UserName,&u.Password)
	if err!=nil{
		//fmt.Println("账号错误")
		//return ""
	}
	if u.Password == password {
		fmt.Println("登录成功")
	} else{
		fmt.Println("密码错误")
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Status: 10000,
		Info: "success",
        RefreshToken: "refreshToken",
		Token: "token",
	})
	freshtoken.UserName = u.UserName
	freshtoken.Password = u.Password
	return string(rs)
}