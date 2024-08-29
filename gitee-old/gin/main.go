package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// binding:"required" 表示必填
	User string `form:"user" json:"user" uri:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
}

func main(){
	r := gin.Default()
   /*
	r.GET("/:name/*action",func(c *gin.Context){
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK,name +" is " +action)
	})

	r.GET("/user",func(c *gin.Context){
		name := c.DefaultQuery("name","默认值")
		c.String(http.StatusOK,name)
	})
	r.POST("/form",func(c *gin.Context){
		username := c.DefaultPostForm("username","默认值")
		pwd := c.PostForm("pwd")
		c.String(http.StatusOK,username +" is " +pwd)
	})
	r.POST("/upload",func(c *gin.Context){
		file,_ := c.FormFile("file")
		c.SaveUploadedFile(file,file.Filename)
		c.String(http.StatusOK,file.Filename)
	})
	*/
	// json 绑定
/*	r.POST("/loginJSON",func(c *gin.Context){
		// 声明接受的变量
		var json Login
		if err := c.ShouldBindJSON(&json);err != nil{
			// 返回错误信息
			// gin.H 封装了生成JSON数据的工具
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		// 判断用户名和密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})*/
	// form 绑定
// 	curl http://localhost:8080/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\",\"password\":\"admin\"}" -X POST
// {"status":"200"}
		/*r.POST("/loginForm",func(c *gin.Context){
		// 声明接受的变量
		var form Login
		if err := c.Bind(&form);err != nil{
			// 返回错误信息
			// gin.H 封装了生成JSON数据的工具
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		// 判断用户名和密码是否正确
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})*/
	// uri绑定
	// curl http://localhost:8080/root/admin
/*	r.GET("/:user/:password",func(c *gin.Context){
		// 声明接受的变量
		var login Login
		if err := c.ShouldBindUri(&login);err != nil{
			// 返回错误信息
			// gin.H 封装了生成JSON数据的工具
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		// 判断用户名和密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})*/
	v1 := r.Group("v1")
	{
		v1.GET("/:name/*action",test)
		v1.GET("/user",test1)

	}
	// r.MaxMultipartMemory = 8 << 20  //限制表单上传的大小 8M
	v2 := r.Group("v2")
	{
		v2.POST("/form",login)
		v2.POST("/upload",upload)

	}
	r.Run()

}
func test( c *gin.Context){
	name := c.Param("name")
	action := c.Param("action")
	c.String(http.StatusOK,name +" is " +action)
}
func test1( c *gin.Context){
	name := c.DefaultQuery("name","默认值")
	c.String(http.StatusOK,name)
}
func login(c *gin.Context){
	username := c.DefaultPostForm("username","默认值")
	pwd := c.PostForm("pwd")
	c.String(http.StatusOK,username +" is " +pwd)
}
func upload(c *gin.Context){
	file,_ := c.FormFile("file")
	c.SaveUploadedFile(file,file.Filename)
	c.String(http.StatusOK,file.Filename)
	// 多个图片
	// form, _:= c.MultipartForm()
	// files :=form.File["files"]
	// for _,file := range files{
	// 	c.SaveUploadedFile(file,file.Filename)
	// }
}