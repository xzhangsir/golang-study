package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	name,city string
	age int8
}
// 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func (p person) eat(){
	fmt.Printf("%v在吃东西",p.name)
}
// 指针类型的接受者
func (p *person) SetAge(newAge int8){
	p.age = newAge
}
// 值类型的接受者
func (p person) SetAge2(newAge int8){
	p.age = newAge
}

// 结构体嵌套
type address struct{
	provinse string
	city string
}
type user struct{
	name string
	gender string
	address address
}
// func main() {		
// 	// p1 := newPerson("张三", "沙河", 90)
// 	// fmt.Printf("%#v\n", p1) //&main.person{name:"张三", city:"沙河", age:90}
// 	// p2 := newPerson("李四", "北京", 99)
// 	// fmt.Printf("%#v\n", p2) //&main.person{name:"李四", city:"北京", age:99}

// 	// Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）
// 	// 函数不需要任何类型 方法属于特定的类型
// 	// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
// 	// 		函数体
// 	// }
// 	// p1 := newPerson("张三", "沙河", 90)
// 	// p1.eat() //张三在吃东西
//   // fmt.Println(p1.age) //90

// 	// //  指针类型的接受者
// 	// // p1.SetAge(30)
// 	// // fmt.Println(p1.age) //30

// 	// // 值类型的接受者 无法修改接收者变量本身
//   // p1.SetAge2(30) // (*p1).SetAge2(30)
// 	// fmt.Println(p1.age) // 90


// 	// 结构体嵌套
// 	// user := user{
// 	// 	name: "小王",
// 	// 	gender:"男",
// 	// 	address:address{
// 	// 		provinse:"山东",
// 	// 		city:"威海",
// 	// 	},
// 	// }
// 	// fmt.Println(user) //{小王 男 {山东 威海}}

// }


// 结构体的继承
// type animal struct{
// 	name string
// }
// func (a *animal) move(){
// 	fmt.Printf("%s会动",a.name)
// }
// type dog struct{
// 	feet int8
// 	*animal
// }
// func (d *dog) wang() {
// 	fmt.Printf("%s会汪汪汪~\n", d.name)
// }
// func main() {
// 	d := dog{
// 		feet:4,
// 		animal: &animal{
// 			name:"李华",
// 		},
// 	}
// 	d.wang() //李华会汪汪汪~
// 	d.move() //李华会动
// }


// 结构体与json序列化
type student struct{
	ID int      `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string
	Name string
}
type class struct{
	Title string `json:"title"`
	Students []*student
}

func main(){
	c := &class{
		Title:"110",
		Students:make([]*student,0,200),
	}
	for i := 0 ; i < 10 ; i++{
		stu := &student{
			Name:fmt.Sprintf("stu%02d",i),
			Gender: "男",
			ID:i,
		}
		c.Students = append(c.Students,stu)
	}
	// fmt.Println(c)
	// //JSON序列化：结构体-->JSON格式的字符串
	data,err := json.Marshal(c)
	if err != nil{
		fmt.Println("序列化失败")
		return
	}
	fmt.Printf("json:%s\n", data)
	// //JSON反序列化：JSON格式的字符串-->结构体
	// str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"}]}`
	// c1 := &class{}
	// err := json.Unmarshal([]byte(str), c1)
	// if err != nil {
	// 	fmt.Println("序列化失败2")
	// 	return
	// }
	// fmt.Printf("%#v\n", c1)
}