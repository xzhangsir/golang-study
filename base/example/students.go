package example

import (
	"fmt"
	"os"
)

type students struct {
	name string
	age  int
	id   int
}
type classes struct {
	allStudens map[int]*students
}

func newStudent(id, age int, name string) *students {
	return &students{
		name: name,
		age:  age,
		id:   id,
	}
}
func (c classes) add() {
	var (
		id   int
		age  int
		name string
	)
	fmt.Println("输入ID")
	fmt.Scan(&id)
	fmt.Println("输入Age")
	fmt.Scan(&age)
	fmt.Println("输入Name")
	fmt.Scan(&name)
	c.allStudens[id] = newStudent(id, age, name)
}
func (c classes) show() {
	for _, v := range c.allStudens {
		fmt.Printf("ID:%v NAME:%v\n", v.id, v.name)
	}
}
func (c classes) del() {
	var id int
	fmt.Println("输入删除的学生ID")
	fmt.Scan(&id)
	_, ok := c.allStudens[id]
	if !ok {
		fmt.Println("没有这个学生")
		return
	}
	delete(c.allStudens, id)
}

func (c classes) edit() {
	var (
		id   int
		age  int
		name string
	)
	fmt.Println("输入编辑的学生ID")
	fmt.Scan(&id)
	v, ok := c.allStudens[id]
	if !ok {
		fmt.Println("没有这个学生")
		c.edit()
		return
	}
	fmt.Printf("%#v", v)
	fmt.Println("输入编辑的学生NAME")
	fmt.Scan(&name)
	fmt.Println("输入编辑的学生age")
	fmt.Scan(&age)
	c.allStudens[id].name = name
	c.allStudens[id].age = age
}

func Students() {
	c := classes{
		allStudens: make(map[int]*students, 50),
	}
	for {
		fmt.Println(`
	        1 ： 查看所有学生
	        2 ： 添加学生
	        3 ： 删除学生
	        4 ： 编辑
					5 ： 按照年龄排序
	        ~~~任意键退出~~~
		`)
		var status int
		fmt.Scan(&status)
		switch status {
		case 1:
			c.show()
		case 2:
			c.add()
		case 3:
			c.del()
		case 4:
			c.edit()
		default:
			os.Exit(0)
		}
	}
}
