package main

import (
	"fmt"
	// "encoding/json"
	"reflect"
)
// func add(x ...int) int {
// 	num := 0
// 	for _,v := range x{
// 		num+=v
// 	}
// 	return num
// }

// func cacl(x,y int,op func(...int) int) int{
// 	return op(x,y)
// }

// type animal struct{
// 	name string
// }
// type dog struct {
// 	age int
// 	animal  //结构体嵌套
// }
func main(){
	// test := cacl(2,3,add)
	// fmt.Print(test)

	// a1 := [...]int{1,4,7,3,8,0,3}
	// a1[0] = 999
	// a2 := a1[1:5:5]

	// fmt.Print(a2)
	// fmt.Print(len(a2),cap(a2))
	// a2 := a1[2:4]
	// a3 := []int{}
	// copy(a2,a3)
	// fmt.Print(len(a2),cap(a2))
	// fmt.Printf("%T\n",a2)	
	// fmt.Print(a3)
	// a4 := make([]bool ,2,10)
	// a4 = append(a4,[]bool{true,false}...)
    // a5 := append(a4[:1],a4[2:]...)
	// fmt.Print(a4,len(a4),cap(a4))
	// fmt.Print(a5)
	// a6 := make([]bool,1,1)
	// copy(a6,a5)
	// fmt.Print(a6)

	// var a = make([]string, 5, 10)
	// fmt.Print(a)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, string(i))
	// }
	// fmt.Println(a)

	// var a7 []int
	// a7 = append(a7,1,4,2,3,5,23,2,7)
	// a8 := []int{90,80,67}
	// copy(a7,a8)
	// fmt.Print(a8)


	// s1 := make(map[string]int,8)
	// s1["张三"] = 12
	// s1["李四"] = 10
	// for key,val := range s1{
	// 	fmt.Print(key,val)
	// }
    
    // s1 := map[string]string{
    // 	"name":"zx",
    // 	"age":"12",
    // }
    // for key,val := range s1{
    // 	fmt.Print(key,val)
    // }
    // _,ok := s1["name"]
    // if ok {
    // 	delete(s1,"name")
    // }
    // fmt.Print(s1)

    // var str1 string = "how do you do"
    // str2 := strings.Split(str1, " ")
    // m1 :=make(map[string]int,20)
    // for _,s := range str2{
    // 	m1[s] = m1[s] + 1
    	// _,ok := m1[s]
    	// if ok {
    	// 	m1[s] += 1
    	// }else{
    	// 	m1[s] = 1
    	// }
    // }
    // fmt.Print(m1)

  
  // num := add(1,4)
  // fmt.Print(num)

  // fmt.Println("1")
  // defer fmt.Println("5")
  // defer fmt.Println("3")
  // fmt.Println(2)
  
   // add := func(x,y int) int{
   //  	return x + y
   //  }
   //  fmt.Print(add(1,9))

  // num := 10
  // p := &num
  // *p = 11
  // fmt.Println(num)
  // fmt.Println(*p)

  // var a = new(int)
  // fmt.Print(*a)
  

// type person struct {
// 	name string
// 	age int
// 	city string
// }
// // var p1 = person{
// // 	"zx",
// // 	12,
// //  "xian"
// // }
// var p1 = person{
// 	name:"zx",
// 	age:12,
// }

// change1 := func(p person){
// 	p.name = "waa"
// }
// change1(p1)
// fmt.Printf("%+v",p1)


// change2 := func(p *person){
// 	p.name = "lq"
// }
// change2(&p1)
// fmt.Printf("%+v",p1)


// 构造函数
// newPerson := func(name string,age int,city string) *person{
// 	return &person{
// 		name:name,
// 		age:age,
// 		city:city,
// 	}
// }

// fmt.Print(newPerson("waa",12,"baoji"))

// var per1 person
// per1.name = "zx"
// per1.city = map[string]int{"xian":12}
// fmt.Print(per1.city["xian"])

// var per2 = new(person)
// fmt.Print(per2)

// per3 := person{
// 	name:"zx",
// 	age:12,
// 	// city:map[string]int{"xian",12}
// }
// fmt.Print(per3)



	// newDog := func(age int,animal animal) dog{
	// 	return dog{
	// 		age:age,
	// 		animal:animal,
	// 	}
	// }
	// d1 := newDog(12,animal{name:"泰迪",})
	// fmt.Printf("%+v",d1)
	// d1.eat()
	// d1.addAge()
	// fmt.Printf("%+v",d1)

	// var a myInt = 9
	// a.hello()
	

	// 序列号和反序列化
	// goJSON()

	// 接口
	// goInterface()
	// 类型断言
	// assign(12)
   
   // 反射  reflect
	// var a int64 = 235
	// testReflect(a)

}

// 方法是作用于特定类型的函数
// 只有dog这个类型 才可以调用eat这个方法
// 值类型接受者
 // func (d dog) eat(){
 // 	fmt.Printf("%s:吃",d.name)
 // }
 // // 指针类型接受者
 // func (d *dog) addAge(){
 // 	d.age++
 // }



// type myInt int

// func (m myInt) hello(){
// 	fmt.Print("xixixi")
// }

// func goJSON(){
//     type students struct{
//     	Id int64 `json:"id"`
//     	Name string `json:"name"`
//     }

//     stu1 := students{
//     	Id:1,
//     	Name:"zx",
//     }

//     // 序列化
//     v,err := json.Marshal(stu1)
//     if err != nil {
//     	fmt.Printf("%v",err)
//     	return
//     }
//     fmt.Printf("%v\n",string(v))
    
//     // 反序列化

//     str := `{"id":12,"name":"zx"}`
//     var stu2  students
//     json.Unmarshal([]byte(str),&stu2)
//     fmt.Printf("%#v\n",stu2)

// }

// 	type moveer interface{
//        move()
// 	}
// 	type animal struct {
// 		name string
// 	}
// 	type person struct {
// 		name string
// 	}
// 	func move(x moveer){
// 		x.move()
// 	}
// 	func (a animal) move(){
// 		fmt.Printf("%v:爬行",a.name)
// 	}
// 	func (p person) move(){
// 		fmt.Printf("%v:走路",p.name)
// 	}
// func goInterface(){
//     var a = animal{
//     	name:"dog",
//     }
//     var p = person{
//     	name:"xiaomin",
//     }
//     move(a)
//     move(p)
// }

// func assign(a interface{}){  //空接口可以接受任意数据类型
// 	v,ok := a.(string)  //类型断言  
// 	if !ok{
// 		fmt.Print("不是string类型")
// 		return
// 	}
// 	fmt.Print(v)
// }

// func testReflect(x interface{}){
//   tp := reflect.TypeOf(x)
//   v := reflect.ValueOf(x)
//   fmt.Println(tp)
//   fmt.Println(v)
// }