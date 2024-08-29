package main

import (
	"fmt"
	"os"
)

type students struct{
	id int64
	name string
}
type classes struct{
	allStudens map[int64]*students
}
func newStudents(id int64,name string)*students{
	return &students{
		id:id,
		name:name,
	}
}
func main(){
	c1 := classes{
		allStudens:make(map[int64]*students,50),
	}
	for{
		fmt.Println(`
	        1 ： 查看所有学生
	        2 ： 添加学生
	        3 ： 删除学生
	        4 ： 编辑
	        ~~~任意键退出~~~
		`)
		var status int
		fmt.Scan(&status)
		switch status {
			case 1 :
				c1.show()
			case 2 :
				c1.add()
			case 3 :
				c1.del()
			case 4 :
				c1.edit()
	        default:
	        	os.Exit(0)
		}
	}
}
func (c classes) show(){
	for _,v := range c.allStudens{
		fmt.Printf("ID:%v NAME:%v\n",v.id,v.name)
	}
}
func (c classes) add(){
	var (
		id int64
		name string
	)
	fmt.Println("输入ID")
	fmt.Scan(&id)
	fmt.Println("输入Name")
	fmt.Scan(&name)
	c.allStudens[id] = newStudents(id,name)
}
func (c classes) del(){
	var id int64
	fmt.Println("输入删除的学生ID")
	fmt.Scan(&id)
	_,ok := c.allStudens[id]
	if !ok {
		fmt.Println("没有这个学生")
		c.del()
		return 
	}
	delete(c.allStudens,id)
}
func (c classes) edit(){
	var (
		id int64
		name string
	)
	fmt.Println("输入编辑的学生ID")
	fmt.Scan(&id)
	v,ok := c.allStudens[id]
	if !ok {
		fmt.Println("没有这个学生")
		c.edit()
		return 
	}
	fmt.Printf("%#v",v)
	fmt.Println("输入编辑的学生NAME")
	fmt.Scan(&name)
	c.allStudens[id].name = name
}