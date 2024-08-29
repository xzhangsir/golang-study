package repository
import (
	// "fmt"
	"ginweb/dao"
)
type Menu struct{
	Id int  `json:"id"`
	Key string  `json:"key"`
	Title string  `json:"title"`
	Path string  `json:"path"`
	ParentId int  `json:"parentId"`
	Auth int `json:"auth"`
}

func GetMenusList()(menus []Menu){
	sqlStr := "select  * from menu"
	rows,_ := dao.DB.Query(sqlStr)
	defer rows.Close()
	for rows.Next(){
		m  := Menu{}
		rows.Scan(&m.Id, &m.Key, &m.Title, &m.Path, &m.ParentId,&m.Auth)
		menus = append(menus,m)
	}
	return
}
func GetMenus(auth string)(menus []Menu){
	sqlStr := "select  * from menu where auth <= ?"
	rows,_ := dao.DB.Query(sqlStr,auth)
	defer rows.Close()
	for rows.Next(){
		m  := Menu{}
		rows.Scan(&m.Id, &m.Key, &m.Title, &m.Path, &m.ParentId,&m.Auth)
		menus = append(menus,m)
	}
	return
}