package main;
import (
	"ginweb/dao"
	"ginweb/controller"
	"github.com/gin-gonic/gin"
	"ginweb/middlewares"
)


func main(){
	dao.InitDB()
	defer dao.Close()
	router := gin.Default()
	router.Use(middlewares.Cors())
	router.GET("/getUserList",controller.GetUserList)
	// id
	router.GET("/getUser",controller.GetUser)
	// username age
	router.POST("/setUser",controller.SetUser)
	// id
	router.POST("/delUser",controller.DelUser)
	//username pwd
	router.POST("/login",controller.Login)
	//usertype
	router.POST("/getMenus",controller.GetMenus)

	router.GET("/getMenuList",controller.GetMenusList)

	v1 := router.Group("/v1")
	{
		v1.GET("/getUserList",controller.GetUserList)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/getMenuList",controller.GetUserList)
	}
	router.Run(":8081")
}