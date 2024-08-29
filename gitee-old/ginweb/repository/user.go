package repository
import (
	// "fmt"
	"ginweb/dao"
)
type User struct{
	Id int  `json:"id"`
	Username string  `json:"username"`
	Age int  `json:"age"`
	Usertype int  `json:"usertype"`
}

func GetUserList()(us []User){
	sqlStr := "select  id,username,age,usertype  from user"
	rows,_ := dao.DB.Query(sqlStr)
	defer rows.Close()
	for rows.Next(){
		u  := User{}
		rows.Scan(&u.Id, &u.Username, &u.Age, &u.Usertype)
		us = append(us,u)
	}
	return
}
func GetUser(id string)(us []User){
	sqlStr := "select  * from user where id = ?"
	u  := User{}
	err := dao.DB.QueryRow(sqlStr,id).Scan(&u.Id, &u.Username, &u.Age)
	if err != nil{
		return nil
	}
	us = append(us,u)
	return
}
func SetUser(username string ,age int)(theID int64){
	// fmt.Println(username,age)
	sqlStr := "insert into user(username,age) values (?,?)"
	ret,_ := dao.DB.Exec(sqlStr,username,age)
	theID, _ = ret.LastInsertId() // 新插入数据的id
	return 
}
func DelUser(id string)(row int64){
	sqlStr := "delete from user where id = ?"
	ret,_ := dao.DB.Exec(sqlStr,id)
	row,_ = ret.RowsAffected()
	return row
}
func Login(username string ,pwd string)(us []User){
	sqlStr := "select id, username,age,usertype from user where username = ? and pwd = ?"
	u  := User{}
	err := dao.DB.QueryRow(sqlStr,username,pwd).Scan(&u.Id, &u.Username,&u.Age, &u.Usertype)
	if err != nil{
		return nil
	}
	us = append(us,u)
	return
}