package controller

import (
	"net/http"
	"ginweb/repository"
	"github.com/gin-gonic/gin"
	// "encoding/json"
	// "fmt"
	"strconv"
	// "reflect"
)

// type User struct{
// 	Id int  `json:"id"`
// 	Username string  `json:"username"`
// 	Age int  `json:"age"`
// }

func GetUserList(c *gin.Context){
	us := repository.GetUserList()
	c.JSON(http.StatusOK,gin.H{
		"status":"200",
		"user":us,
	})
}
func GetUser(c *gin.Context){
	id := c.Query("id")
	us := repository.GetUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status":"200",
		"user":us,
	})
}
func SetUser(c *gin.Context){
	username := c.PostForm("username")
	age,_ := strconv.Atoi(c.PostForm("age"))
	if username == ""{
		c.JSON(http.StatusOK,gin.H{
			"status":"404",
			"msg":"username不能为空",
		})
		return
	}
	if age == 0{
		c.JSON(http.StatusOK,gin.H{
			"status":"404",
			"msg":"age必须大于0",
		})
		return
	}
	id := repository.SetUser(username,age)
	c.JSON(http.StatusOK,gin.H{
		"status":"200",
		"id":id,
	})
}
func DelUser(c *gin.Context){
		id := c.PostForm("id")
		row :=  repository.DelUser(id)
		if row == 0{
			c.JSON(http.StatusOK,gin.H{
				"status":"404",
				"msg":"未找到记录",
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":"200",
				"msg":"删除成功",
			})
		}
}
func Login(c *gin.Context){
	username := c.PostForm("username")
	pwd := c.PostForm("pwd")
	us :=  repository.Login(username,pwd)
    c.JSON(http.StatusOK,gin.H{
		"status":"200",
		"user":us,
	})
}


// func UserId(c *gin.Context){
// 	b,_ := c.GetRawData()
// 	// 定义map或结构体
// 	var m map[string]interface{}
// 	// 反序列化
// 	_ = json.Unmarshal(b, &m)
// 	// fmt.Print(m)

// 	c.JSON(http.StatusOK, m)
// }