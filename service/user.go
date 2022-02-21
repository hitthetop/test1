package service

import (
	"encoding/json"
	"fmt"
	"myproject/dao"
	"myproject/function"
	"myproject/model"
)

func ChangePassword(password string,newPassword string) string {
	var u model.User
	sqlStr := "select name,password from user where name = ?"
	rowOBJ:= dao.DB.QueryRow(sqlStr,freshtoken.UserName)
	err := rowOBJ.Scan(&u.UserName,&u.Password)
	if err!=nil{
		fmt.Println("账号或密码错误1")
		return ""
	}
	if u.Password == password {
		sql := "UPDATE user SET password = ? WHERE name = %s"
		str := fmt.Sprintf(sql, freshtoken.UserName)
		_, err := dao.DB.Exec(str, newPassword)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("修改成功")
	}
	rs , err := json.Marshal(model.Data{
		Status: 10000,
		Info: "success",
	})
	freshtoken.UserName = u.UserName
	freshtoken.Password = newPassword
	return string(rs)
}
func QueryInformation(username string) string {
	var u model.User
	sqlstr := "select introduction,emil,phone,qq,gender,birth from user where name = ?"
	rowOBJ:= dao.DB.QueryRow(sqlstr,freshtoken.UserName)
	err := rowOBJ.Scan(&u.Introduction,&u.Emil,&u.Phone,&u.QQ,&u.Gender,&u.Birth)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.User{
		UserName: freshtoken.UserName,
		Introduction: u.Introduction,
		Emil: u.Emil,
		Phone: u.Phone,
		QQ: u.QQ,
		Gender: u.Gender,
		Birth: u.Birth,
	})
	return string(rs)
}
func ChangeInformation(name string,emil string,introduction string,phone int,qq int,
	gender string,birth string) string {
	sqlStr:="UPDATE user SET " +
		"name = ?,introduction= ?,emil= ?," +
		"phone= ?,qq= ?,gender= ?,birth= ? WHERE name = ?"
	_,err :=dao.DB.Exec(sqlStr,name,introduction,emil,phone,qq,gender,birth,freshtoken.UserName)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.User{
		UserName: name,
		Introduction: introduction,
		Emil: emil,
		Phone: phone,
		QQ: qq,
		Gender: gender,
		Birth: birth,
	})
	return string(rs)
}

// CollectTopic 收藏话题
func CollectTopic(id string) string {
	sql := `insert into %s(topicName)values(?)`
	str := fmt.Sprintf(sql,freshtoken.UserName)
	_,err :=dao.DB.Exec(str,id)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}

// FollowUser 关注用户
func FollowUser(id string) string {
	sql := `insert into %s(name)values(?)`
	str := fmt.Sprintf(sql,freshtoken.UserName)
	_,err :=dao.DB.Exec(str,id)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}
func Nobody() string {
	freshtoken.UserName="匿名"
	freshtoken.Password="匿名"
	rs , _ := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}

// GetCollection 获取搜藏列表
func GetCollection() string {
	var u model.Collection
	var N []byte
	sql := "select topicName from %s "
	str := fmt.Sprintf(sql,freshtoken.UserName)
	rows,err:=dao.DB.Query(str)
	if err != nil{
		return ""
	}
	for rows.Next(){
		err := rows.Scan(&u.TopicName)
		if err!=nil{
			fmt.Println(err)
			return ""
		}
		rs , err := json.Marshal(model.Collection{
			TopicName: u.TopicName,
		})
		N = function.BytesCombine(N,rs)
	}
	rows.Close()
	return string(N[:])
}

// GetPeople 获取关注列表
func GetPeople() string {
	var u model.Collection
	var N []byte
	sql := "select topicName from %s "
	str := fmt.Sprintf(sql,freshtoken.UserName)
	rows,err:=dao.DB.Query(str)
	if err != nil{
		return ""
	}
	for rows.Next(){
		err := rows.Scan(&u.Name)
		if err!=nil{
			fmt.Println(err)
			return ""
		}
		rs , err := json.Marshal(model.Collection{
			Name: u.Name,
		})
		N = function.BytesCombine(N,rs)
	}
	rows.Close()
	return string(N[:])
}
func DislikePeople(name string) string {
	sql := `insert into %s(dislikeName)values(?)`
	str := fmt.Sprintf(sql,freshtoken.UserName)
	_,err :=dao.DB.Exec(str,name)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}