package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)
func main()  {
	initDB()//make sure the connection is well
    menu()//use menu to choose what you want to do
}
type userinfo struct{
	Name string
	password string
}
var user struct{
	Name string `db:"id"`
	Password string `db:"password"`
}
var security struct{
	Name string `db:"name"`
	Questions string `db:"questions"`
	Answers string `db:"answers"`
}
var Db *sqlx.DB//Db数据库连接池
func initDB() (err error){
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/users")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	err = Db.Ping()
	if err != nil{
		fmt.Printf("open %s faild, err:%v\n",err)
		return
	}
	fmt.Println("连接成功！")
	return
}
func menu(){
	fmt.Println("1.login\n","2.register\n","3.changepassword\n","4.exit")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		login()
		return
	case 2:
		register()
		return
	case 3:
		changepassword()
		return
	case 4:
		exit()
		break
	default:
		fmt.Println("please write down the right option")
		return
	}
}
var scanner  *bufio.Scanner
func getInput() string {
	scanner.Scan()
	return scanner.Text()
}//用函数接收界面开始时输入的值
func login()  {
	fmt.Print("请输入您的账号：")
	fmt.Scanln(&user.Name)
	fmt.Print("请输入您的密码：")
	fmt.Scanln(&user.Password)
	var u userinfo
	sqlStr := "select name,password from user where name = ?"
	rowOBJ:=Db.QueryRow(sqlStr,user.Name)
	err:=rowOBJ.Scan(&u.Name,&u.password)
	if err!=nil{
		fmt.Println("账号错误")
	}
    if u.password == user.Password{
		fmt.Println("恭喜您登录成功")
	} else{
		fmt.Println("密码错误")
	}
}
func register()  {
	fmt.Print("请输入您的账号：")
	fmt.Scanln(&user.Name)
	fmt.Print("请输入您的密码：")
	fmt.Scanln(&user.Password)
	r := "insert into user(Name,Password) values (?,?)"
	re,err := Db.Exec(r,user.Name,user.Password)
	if err != nil{
		fmt.Printf("insert failed,err:%v\n",err)
		return
	}
	newID,err :=re.LastInsertId()
	if err != nil{
		fmt.Printf("get lastinsert id failed,err:%v\n",err)
	}
	fmt.Printf("insert success, the id is %d.\n", newID)
	fmt.Println("congratulations！the ID have changed successfully")
}
func changepassword() {
	fmt.Print("请输入您的账号：")
	fmt.Scanln(&user.Name)
	fmt.Print("请输入您的密码：")
	fmt.Scanln(&user.Password)
	var u userinfo
	sqlStr := "select name,password from user where name = ?"
	rowOBJ:=Db.QueryRow(sqlStr,user.Name)
	err:=rowOBJ.Scan(&u.Name,&u.password)
	if err!=nil{
		fmt.Println("账号错误")
	}
	if u.password == user.Password{
		fmt.Println("恭喜您登录成功")
	} else{
		fmt.Println("密码错误")
		os.Exit(0)
	}
	fmt.Print("请输入您的账号：")
	fmt.Scanln(&security.Name)
	fmt.Print("请输入您的密保问题：")
	fmt.Scanln(&security.Questions)
	fmt.Print("请输入您的密保答案：")
	fmt.Scanln(&security.Answers)

	str := "insert into security(name,question,answer) values (?,?,?)"

	re, err := Db.Exec(str, security.Name, security.Questions, security.Answers)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	new, err := re.LastInsertId()
	if err != nil {
		fmt.Printf("set 密保 failed,err:%v\n", err)
	}
	fmt.Printf("insert success, the id is %d.\n", new)
	fmt.Println("恭喜您，密保设置成功")
}
func exit()  {
fmt.Println("good bye")
}