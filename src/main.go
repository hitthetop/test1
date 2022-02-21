package main

import (
	"myproject/dao"
	"myproject/total"
)

func main()  {
	dao.InitDB()
	total.Menu()
}
