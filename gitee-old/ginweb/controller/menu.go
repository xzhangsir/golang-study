package controller

import (
	"net/http"
	"ginweb/repository"
	"github.com/gin-gonic/gin"
	// "fmt"
	// "reflect"
	// "encoding/json"
)
func GetMenusList(c *gin.Context){
	menus := repository.GetMenusList()
	c.JSON(http.StatusOK,gin.H{
		"status":"200",
		"menus":menus,
	})
}
func GetMenus(c *gin.Context){
	usertype := c.PostForm("usertype")
	menus := repository.GetMenus(usertype)
	c.JSON(http.StatusOK,gin.H{
		"status":"200",
		"menus":menus,
	})
}

