package total

import (
	"github.com/gin-gonic/gin"
	"myproject/function"
	"myproject/service"
	"net/http"
)

func Menu()  {
	r := gin.Default()
	//匿名登陆实现
	//需要先在数据库中创建一个名称为匿名，密码为匿名，所有字段值均为匿名的数据
	r.GET("/nobody", func(c *gin.Context) {
		B := service.Nobody()
		if B==""{
			c.String(http.StatusFailedDependency,"匿名成功")
		}
		if B!=""{
			c.String(http.StatusOK,B)
		}
	})
	//登陆接口
	r.GET("/user/token", func(c *gin.Context) {
		name := c.Query("name")
		password := c.Query("password")
		B := service.Login(name,password)
		if B==""{
			c.String(http.StatusFailedDependency,"登陆失败")
		}
		if B!=""{
			c.String(http.StatusOK,B)
		}
	})
	//修改密码接口
	r.PUT("/user/password", func(c *gin.Context) {
		newPassword:=c.PostForm("newPassword")
		oldPassword:=c.PostForm("oldPassword")
		B := service.ChangePassword(oldPassword,newPassword)
		if B == ""{
			c.String(http.StatusFailedDependency,"修改失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//拉黑用户
	r.GET("/dislike", func(c *gin.Context) {
		name := c.Query("name")
		B := service.DislikePeople(name)
		if B == ""{
			c.String(http.StatusFailedDependency,"拉黑失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//注册接口
	r.POST("/user/register", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		introduction := c.PostForm("introduction")
		emil := c.PostForm("emil")
		phone := c.GetInt("phone")
		qq := c.GetInt("qq")
		gender := c.PostForm("gender")
		birth := c.PostForm("birth")
		B := service.Register(name,password,emil,introduction,phone,qq,gender,birth)
		if B==""{
			c.String(http.StatusFailedDependency,"注册失败")
		}
		if B!=""{
			c.String(http.StatusOK,B)
		}
	})
	//刷新token
	r.GET("/user/token/refresh", func(c *gin.Context) {
   B :=  service.Token()
		if B == ""{
			c.String(http.StatusFailedDependency,"刷新失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//查询资料接口
	r.GET("/user/info/{user_id}", func(c *gin.Context) {
		name := c.Query("user_id")
		B := service.QueryInformation(name)
		if B == ""{
			c.String(http.StatusFailedDependency,"查询失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//修改资料接口
	r.PUT("/user/info", func(c *gin.Context) {
		nickname := c.PostForm("nickname")
		introduction := c.PostForm("introduction")
		emil := c.PostForm("emil")
		phone := c.GetInt("phone")
		qq := c.GetInt("qq")
		gender := c.PostForm("gender")
		birth := c.PostForm("birth")
		B:=service.ChangeInformation(nickname,emil,introduction,phone,qq,gender,birth)
		if B == ""{
			c.String(http.StatusFailedDependency,"修改失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//发起话题接口
	r.POST("/post/single", func(c *gin.Context) {
		content := c.PostForm("content")
		topicName := c.PostForm("topicName")
		B := service.InsertTopic(content,topicName)
		if B==false{
			c.String(http.StatusFailedDependency,"加入话题失败")
		}
		if B==true{
			c.String(http.StatusOK,"加入话题成功")
		}
	})
	//展示所有话题
	r.GET("/topic/list", func(c *gin.Context) {
    v := function.Ioutil("topicName.txt")
	c.String(http.StatusOK,v)
	})
	//展示话题接口
	r.GET("/post/single/{post_id}", func(c *gin.Context) {
		topicName := c.Query("topicName")
		B := service.ShowTopic(topicName)
		if B == ""{
			c.String(http.StatusFailedDependency,"未找到标题")
		}
		if B!=""{
			c.String(http.StatusOK,B)
		}
	})
	//评论话题接口
	r.GET("/comment", func(c *gin.Context) {
		topicName := c.Query("topicName")
		message := c.Query("message")
		B := service.ShowTopic(topicName)
		if B == ""{
			c.String(http.StatusFailedDependency,"未找到标题")
		}
		if B!=""{
			c.String(http.StatusOK,B)
		}
		C := service.InsertComment(topicName,message)
		if C==""{
			c.String(http.StatusFailedDependency,"未找到标题")
		}else {
			c.String(http.StatusOK,C)
		}
	})
	//更新话题
	r.PUT("/post/single/{post_id}", func(c *gin.Context) {
		name:=c.PostForm("topicName")
		content:=c.PostForm("content")
		title := c.PostForm("title")
    C := service.ChangeTopicName(name,content,title)
		if C=="" {
			c.String(http.StatusFailedDependency,"修改失败")
		}
		if C!=""{
			c.String(http.StatusOK,C)
		}
	})
	//删除文章
	r.DELETE("/post/single/{post_id}", func(c *gin.Context) {
     id := c.Query("post_id")
	 B := service.DeleteTopic(id)
	 if B == ""{
		 c.String(http.StatusFailedDependency,"删除失败")
	 }
	 if B != ""{
		 c.String(http.StatusOK,B)
	 }
	})
	//点赞评论
	r.PUT("/operate/praise", func(c *gin.Context) {
	name := c.PostForm("topicName")
    id := c.PostForm("id")
	B := service.GetLike(name,id)
		if B == ""{
			c.String(http.StatusFailedDependency,"点赞失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//关注用户
	r.PUT("/operate/focus", func(c *gin.Context) {
		id := c.PostForm("user_id")
        B := service.FollowUser(id)
		if B == ""{
			c.String(http.StatusFailedDependency,"关注失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//获取关注列表
	r.GET("/operate/focus/list", func(c *gin.Context) {
		B := service.GetPeople()
		if B == ""{
			c.String(http.StatusFailedDependency,"获取失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//收藏话题
	r.PUT("/operate/collect", func(c *gin.Context) {
		id := c.PostForm("post_id")
		B := service.FollowUser(id)
		if B == ""{
			c.String(http.StatusFailedDependency,"收藏失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//获取用户收藏列表
	r.GET("/operate/collect/list", func(c *gin.Context) {
    B := service.GetCollection()
		if B == ""{
			c.String(http.StatusFailedDependency,"获取失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	r.Run(":8000")
}