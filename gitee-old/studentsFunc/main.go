package main

import (
	"fmt"
	"os"
)

type students struct{
	id int64
	name string
}
var stu = make(map[int64]*students,50)
func newStudents(id int64,name string) *students{
	return &students{
		id:id,
		name:name,
	}
}
func main(){
	
for{
		fmt.Println(`
	        1 ： 查看所有学生
	        2 ： 添加学生
	        3 ： 删除学生
	        4 ： 退出
		`)
		var status int
		fmt.Scan(&status)
		switch status {
			case 1 :
				showAllStudents()
			case 2 :
				addStudents()
			case 3 :
				delStudents()
			case 4 :
				os.Exit(0)
	        default:
	        	fmt.Println("gun")
		}
		fmt.Print(status)
	}
}

func showAllStudents(){
	for _,v := range stu{
		fmt.Printf("学号：%v  姓名：%v", v.id,v.name)
	}
}
func addStudents(){
	var (
		id int64
		name string
	)
	fmt.Println("输入ID")
	fmt.Scan(&id)
	fmt.Println("输入Name")
	fmt.Scan(&name)
	stu[id] = newStudents(id,name)
}
func delStudents(){
	var id int64
	fmt.Println("输入删除的学生ID")
	fmt.Scan(&id)
	delete(stu,id)
}
