package main

import (
	// "fmt"
	// "strings"
	// "math/rand"
  // "time"
	// "sort"
	"hello/main/testpackage"
)

// var 变量名 类型 = 表达式
// var name string = "zx"
// var age int = 18
// 也可
// var name,age = "zx",18

// 常量
// const pi = 3.1415

// const (
// 		n1 = iota //0
// 		n2        //1
// 		n3        //2
// 		n4        //3
// 	)

// const (
// 		n1 = iota //0
// 		n2 = 100  //100
// 		n3 = iota //2
// 		n4        //3
// 	)

// const (
// 		a, b = iota + 1, iota + 2 //1,2
// 		c, d                      //2,3
// 		e, f                      //3,4
// 	)


func main()  {
	// 在函数内部 短变量声明
	// num := 12
	// 匿名变量用一个下划线 _ 表示
	// fmt.Println("hello",n4)
	// var str string = "hello"
	// 长度
	// fmt.Println(len(str)) //5
	// 分割
	// fmt.Println(strings.Split(str,"l")) // [he  o]
	// 拼接
	// fmt.Println(fmt.Sprintf(str + "world")) //helloworld
	// 是否包含
	// fmt.Println(strings.Contains(str,"ll"))  //true
	// 字符串是否已xx开头
	// fmt.Println(strings.HasPrefix(str,"he")) //true
	// 字符串是否已xx结尾
	// fmt.Println(strings.HasSuffix(str,"he")) //false
	// 字符串出现的位置 Index
	// fmt.Println(strings.LastIndex(str,"9")) // -1
	// 将字符串切片按照指定字符连接
	// sce := []string{"2023","02"}
	// fmt.Println(strings.Join(sce,"-")) //2023-02 

	// s1 := "hello语言"
	// for i := 0 ; i < len(s1) ;i++{ //byte
	// 	fmt.Println(i,s1[i])
	// }
	// for _,r := range s1{  //rune
	// 	fmt.Println(r)
	// }
  // 修改字符串
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
	// 无论哪种转换，都会重新分配内存，并复制字节数组。
	// s1 := "big"
	// byteS1 := []byte(s1)
	// byteS1[0] = 'p'
	// fmt.Println(string(byteS1))
	
	// s2 := "大西瓜"
	// runeS2 := []rune(s2)
	// runeS2[0] = '红'
	// fmt.Println(string(runeS2))

	// 汉字出现的个数
	// s3 := "hello沙河小王子"
	// count := 0
	// for _,r := range s3{
	// 	if r > 'z'{
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	// ----------------数组------------------
	// 数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化
	// [n]*T表示指针数组，*[n]T表示数组指针 。

	// var testArray [3]int              //[0 0 0]          //数组会初始化为int类型的零值
	// var numArray = [3]int{1, 2}       //[1 2 0]          //使用指定的初始值完成初始化
	// var cityArray = [3]string{"北京", "上海", "深圳"}  //[北京 上海 深圳]   //使用指定的初始值完成初始化
	// // 自行推断数组的长度
	// var numArray = [...]int{1, 2}     //[1 2]    
	// var cityArray = [...]string{"北京", "上海", "深圳"} //[北京 上海 深圳]  
	// // 使用指定索引值的方式来初始化数组
	// a := [...]int{1: 1, 3: 5}       // [0 1 0 5]

	// var cityArray = [...]string{"北京", "上海", "深圳"}
	
	// for i:= 0 ; i < len(cityArray);i++{
	// 	fmt.Println(cityArray[i])
	// }

	// for k,v := range cityArray{
	// 	fmt.Println(k,v)
	// }
	// ------数组求和-----
	// var num = [...]int{1, 3, 5, 7, 8}
	// var count = 0
	// for i := 0 ; i < len(num) ;i++{
	// 	count += num[i]
	// }
	// fmt.Println(count)
	// ------twoSum---------
	// 从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	// var num = [...]int{1, 3, 5, 7, 8}
	// var count int = 8;
	// for i := 0 ; i < len(num) ; i++{
	// 	var temp = count - num[i]
	// 	for j := i ; j < len(num) ;j++{
	// 		if num[j] == temp{
	// 			fmt.Println(i,j)
	// 		}
	// 	}
	// }

	// ----------切片-----------
	// var num = [...]int{1,3,5,7,9} //数组
	// var s = num[1:3] //切片 左包含 右不包含
	// // len 长度 cap 容量
	// fmt.Println(len(s),cap(s),s,len(num)) //2 4 [3 5] 5

	// 对切片再切片 右边界的上限是切片的容量
	// a := [6]int{1,2,3,4,5,7}
	// s := a[2:4]
	// fmt.Println(s,len(s),cap(s))//[3 4] 2 4
	// s2 := s[:4]  // 索引的上限是cap(s)而不是len(s)
	// fmt.Println(s2,len(s2),cap(s2)) //[3 4 5 7] 4 4

	// 使用make()函数构造切片  make([]T, size, cap)
	// a := make([]int, 2, 6)
	// fmt.Println(a,len(a),cap(a)) //[0 0] 2 6

	// 对切片便利
	// s := []int{1,3,5,7}
	// for i := 0 ; i < len(s) ; i++{
	// 	fmt.Println(s[i],i)
	// }
	// for index,val := range s{
	// 	fmt.Println(index,val)
	// }

  // 切片中的方法
	//  append
  // var s []int
	// fmt.Println(s,len(s),cap(s)) //[] 0 0
	// s = append(s,66,77)
	// fmt.Println(s,len(s),cap(s)) //[66 77] 2 2

	// copy
	// a := []int{1,3,5,7}
	// c := make([]int, 5)
	// copy(c,a)
	// fmt.Println(a) //[1 3 5 7]
	// fmt.Println(c) //[1 3 5 7 0]

	// 从切片中删除元素
	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// a = append(a[:3],a[4:]...) //删除索引为3的元素
	// fmt.Println(a) //[30 31 32 34 35 36 37]

	// 数组排序
	// var a = [...]int{3, 7, 8, 9, 1}
	// for i := 0 ; i < len(a) ; i++{
	// 	for j := i + 1 ; j < len(a) ; j++{
	// 		if(a[i] > a[j]){
	// 			a[i],a[j] = a[j],a[i]
	// 		}
	// 	}
	// }
	// fmt.Println(a) //[1 3 7 8 9]



	// ---------map-----------
	// map类型的变量默认初始值为nil  cap容量不是必须的
	// make(map[keyType]ValueType,[cap])
	// 	scoreMap := make(map[string]int, 2)
	// 	scoreMap["张三"] = 90
	// 	scoreMap["李四"] = 90
	// 	scoreMap["王五"] = 90
	// 	fmt.Println(scoreMap) //map[张三:90 李四:90 王五:90]
	// 	fmt.Printf("type of a:%T\n", scoreMap) //type of a:map[string]int

	// // map也支持在声明的时候填充元素，例如：
		// userInfo := map[string]string{
		// 	"username": "zhangsan",
		// 	"pwd": "123",
		// }
		// 判断map中建是否存在
		// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
		// fmt.Println(userInfo) 
		// value,ok := userInfo["username"]
		// fmt.Println(value,ok)  //zhangsan true
		// value,ok := userInfo["age"]
		// fmt.Println(value,ok)  //  false

		// for key,value := range userInfo{
		// 	fmt.Println(key,value)
		// }
		// 只遍历key
		// for key := range userInfo{
		// 	fmt.Println(key)
		// }
		// // 只遍历value
		// for _,value := range userInfo{
		// 	fmt.Println(value)
		// }

		// 删除一组键值对  delete(map, key)
		// scoreMap := make(map[string]int)
		// scoreMap["张三"] = 10
		// scoreMap["小明"] = 20
		// scoreMap["李华"] = 30
		// delete(scoreMap,"小明")
		// fmt.Println(scoreMap)

    // rand.Seed(time.Now().UnixNano()) //初始化随机数种子
		// var scoreMap = make(map[string]int,10)
		// for i := 0 ; i < 10 ; i++{
		//	// 生成0~99的随机整数
		// 	scoreMap[fmt.Sprintf("str%v" , i)] = rand.Intn(100)
		// }
		// // fmt.Println(scoreMap)
		// var keys = make([]string,0)
		// for key := range scoreMap{
		// 	keys = append(keys,key)
		// }
		// // fmt.Println(keys)
		// // 对切片排序
		// sort.Strings(keys)
		// // fmt.Println(keys)
		// for _,key := range keys{
		// 	fmt.Println(key,scoreMap[key])
		// }

	// 元素为map类型的切片
	// var mapSlice = make([]map[string]string,3)
	// mapSlice[0] = make(map[string]string,10)
	// mapSlice[0]["username"] = "zhangsan"
	// mapSlice[0]["pwd"] = "123"
	// fmt.Println("mapSlice",mapSlice) //[map[pwd:123 username:zhangsan] map[] map[]]

	// 值为切片的map
	// var sliceMap = make(map[string][]string,3)
	// value,ok := sliceMap["info"]
	// if !ok{
	// 	value = make([]string,0,2)
	// }
	// value = append(value,"背景","上海")
	// sliceMap["info"] = value
	// fmt.Println(sliceMap) //map[info:[背景 上海]]

	// 统计字符串中没个字母出现的次数
	// var str = "how do you do"
	// var wordMap = make(map[string]int)

	// words := strings.Split(str,"")
	// // for i := 0 ; i < len(words) ; i++{
	// // 	key := words[i]
	// // 	wordMap[key] += 1
	// // }
	// for _,key := range(words){
	// 	wordMap[key] += 1
	// }
	// fmt.Println(wordMap)


	// 结构体
	// type person struct{
	// 	name,sex string
	// 	age int
	// }
	// var p1 person
	// p1.name = "zx"
	// p1.sex = "nan"
	// p1.age = 12
	// fmt.Println(p1,p1.name) //{zx nan 12} zx

	// 匿名结构体
	// var user struct{name string; age int}
	// user.name = "lisi"
	// user.age = 12
	// fmt.Println(user,user.name) //{lisi 12} lisi

	// 创建指针类型结构体
	// var p2 = new(person)
	// p2.name = "小王子"
	// p2.age = 28
	// p2.sex = "上海"
	// fmt.Println(p2,p2.name) //&{小王子 上海 28} 小王子

	// 取结构体的地址实例化
	// 使用&对结构体进行取地址操作
	// 相当于对该结构体类型进行了一次new实例化操作
	// var p3 = &person{}
	// p3.name = "小王子"
	// p3.age = 28
	// p3.sex = "上海"
	// fmt.Println(p3,p3.name) //&{小王子 上海 28} 小王子

	// 结构体实例化
	// 使用键值对初始化
	// p5 := person{
	// 	name: "小王子",
	// 	sex: "北京",
	// 	age:  18,
	// }
	// fmt.Println(p5,p5.name) //{小王子 北京 18} 小王子

	// 对结构体指针进行键值对初始化
	// p6 := &person{
	// 	name: "小王子",
	// 	sex: "北京",
	// 	age:  18,
	// }
	// fmt.Println(p6,p6.name)  //&{小王子 北京 18} 小王子
	// // 部分字段初始化
	// p7 := &person{
	// 	name: "小王子",
	// }
	// fmt.Println(p7,p7.name) //&{小王子  0} 小王子

	// 结构体占用一块连续的内存。
	// type test struct {
	// 	a int8
	// 	b int8
	// 	c int8
	// 	d int8
	// }
	// n := test{
	// 	1, 2, 3, 4,
	// }
	// fmt.Printf("n.a %p\n", &n.a) //n.a 0xc000016088
	// fmt.Printf("n.b %p\n", &n.b) //n.b 0xc000016089
	// fmt.Printf("n.c %p\n", &n.c) //n.c 0xc00001608a
	// fmt.Printf("n.d %p\n", &n.d) //n.d 0xc00001608b


	// type student struct{
	// 	name string
	// 	age int
	// }

	// m := make(map[string]*student)
	// stus := []student{
	// 	{name: "小王子", age: 18},
	// 	{name: "娜扎", age: 23},
	// 	{name: "大王八", age: 9000},
	// }

	// for _,stu := range stus{
		
	// 	m[stu.name] = &stu
	// 	/**
	// 	由于我们在分配之前创建了一个临时变量 stu，因此所有指针都指向了同一个变量：最后一次迭代的学生*/
	// 	// a := stu
	// 	// m[stu.name] = &a
	// }
	// for key,val := range m{
	// 	fmt.Println(key,val)
	// }


	testpackage.Test()

}


