package base

func Init() {
	// 小练习
	// example.Students()
	// example.ParseIni()
	// example.ObserveFunc()
	// example.EgFunc()
	// 进阶和技巧
	High()
	//并发
	// GoroutineFunc()
	// 命令行参数
	// flagFunc()
	// 文件操作
	// FileFunc()
	/* // ------- 数组 -------
	// 同一种数据类型元素的集合
	// 数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化
	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	// [n]*T表示指针数组，*[n]T表示数组指针 。
	// var arr [2]int
	// arr := [2]string{"a", "b"}
	// arr := [...]string{1: "a", 2: "b"}
	arr := [...]int{2, 3, 4}
	fmt.Println(arr, len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for index, val := range arr {
		fmt.Println(index, val)
	}
	add := func(x []int) int {
		num := 0
		for _, v := range x {
			num += v
		}
		return num
	}

	fmt.Println(add(arr[:])) */

	/* // 多维数组
	// 只支持第一层 ...
	// arr := [...][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	arr := [2][3]int{
		{1, 11, 111},
		{2, 22, 222},
	}
	fmt.Println(arr)
	for _, rows := range arr {
		for _, val := range rows {
			fmt.Printf("%v\t", val)
		}
		fmt.Println()
	} */

	// ------- 切片 []T-------
	// var s []int
	// var s = []int{1, 2, 3}
	/* 	// 追加append
	   // 每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
		//  当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，
		// 此时该切片指向的底层数组就会更换
	   	var s = make([]int, 2, 4) //[0 0] 2 4
	   	s = append(s, 1, 1)       //[0 0 1 1] 4 4
	   	temp := []int{8, 9, 10}
	   	s = append(s, temp...) //[0 0 1 1 8 9 10] 7 8
	   	fmt.Println(s, len(s), cap(s)) */
	/* 	// copy
	   	var s1 = []int{2, 4, 6, 8}
	   	s2 := s1
	   	fmt.Println(s1, s2) //[2 4 6 8] [2 4 6 8]
	   	s2[2] = 44
	   	fmt.Println(s1, s2) //[2 4 44 8] [2 4 44 8]
	   	s3 := make([]int, 5)
	   	copy(s3, s1) //s1拷贝到s3
	   	s3[3] = 33
	   	fmt.Println(s1, s3) //[2 4 44 8] [2 4 44 33 0] */
	/* 	// 移除
	   	// 切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	   	s := []int{1, 3, 5, 7, 9}
	   	s = append(s[:2], s[3:]...) //移除第2个
	   	fmt.Println(s)              //[1 3 7 9] */
	/* s := []int{3, 1, 4, 2, 9}
	sort.Ints(s) //升序
	// sort.Sort(sort.Reverse(sort.IntSlice(s))) l//降序
	// index := sort.SearchInts(s, 4) //s中找5
	// if index < len(s) && s[index] == 4 {
	// 	fmt.Println("找到了")
	// } else {
	// 	fmt.Println("没在")
	// }
	// fmt.Println(s, index)
	// 二分查找
	// index := sort.Search(len(s), func(i int) bool {
	// 	return s[i] >= 4
	// })
	// fmt.Println(index) */

	/* // ---------map----make(map[keyType]ValueType,[cap])-------
	// m := make(map[string]string, 2)
	// m["a"] = "aa"
	// m["c"] = "bb"
	// m := map[string]string{
	// 	"a": "aa",
	// 	"b": "bb",
	// 	"c": "cc",
	// }
	// fmt.Printf("%p %v\n", m, m)
	// value, ok := m["a"]    //判断map中建是否存在
	// fmt.Println(value, ok) //aa  true

	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }
	// // 只遍历key
	// for key := range m {
	// 	fmt.Println(key)
	// }
	// // 只遍历value
	// for _, value := range m {
	// 	fmt.Println(value)
	// }
	// delete(m, "b") //删除一组键值对
	// fmt.Println(m)

	//元素为map类型的切片
	// s := make([]map[string]string, 3)
	// s[0] = map[string]string{
	// 	"a": "aa",
	// 	"b": "bb",
	// }
	// s[1] = make(map[string]string, 2)
	// s[1]["b"] = "bb"
	// s[1]["c"] = "cc"
	// fmt.Println(s)

	// 值为切片的map
	// s := map[string][]int{
	// 	"a": {2, 3, 4},
	// }
	// v, ok := s["a"]
	// if !ok {
	// 	v = make([]int, 0, 2)
	// }
	// v = append(v, 99)
	// s["a"] = v
	// // s["a"] = append(s["a"], 6)
	// fmt.Println(s)

	// 统计字符串中每个单词出现的次数
	// var str = "how do you do"
	// m := make(map[string]int)
	// s := strings.Split(str, " ")
	// for _, v := range s {
	// 	m[v] += 1
	// }
	// fmt.Println(m) */
	/*
		// ---------指针-------
		// &取地址 *根据地址取值
		// a := 10
		// b := &a
		// fmt.Println(*b, a, b)
		// *b = 12
		// fmt.Println(*b, a, b)

		// var a *int
		// a = new(int) //对a进行初始化后 才有内存空间 才可以赋值
		// *a = 100
		// fmt.Println(a)

		// 	new与make的区别
		// 二者都是用来做内存分配的。
		// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
		// 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	*/

	/* 	// --------结构体------------
	   	type person struct {
	   		name, sex string
	   		age       int
	   		address   []string
	   	}
	   	// var p person
	   	// p.name = "zx"
	   	// p.sex = "man"
	   	// p.age = 12
	   	// p.address = []string{"陕西", "西安"}

	   	// p := person{
	   	// 	name:    "zx",
	   	// 	sex:     "man",
	   	// 	age:     12,
	   	// 	address: []string{"陕西", "宝鸡"},
	   	// }

	   	// p := new(person)
	   	// p.name = "zx" //等同于  	(*p).name = "zx"
	   	// fmt.Printf("%p %v\n", p, *p)

	   	// p := &person{ //相当于进行了一次new实例化操作。
	   	// 	name: "zx",
	   	// }
	   	// fmt.Println(*p)

	   	// 构造函数
	   	// newPerson := func(name, sex string, age int, address []string) *person {
	   	// 	return &person{
	   	// 		name:    name,
	   	// 		sex:     sex,
	   	// 		age:     age,
	   	// 		address: address,
	   	// 	}
	   	// }
	   	// p := newPerson("zx", "man", 12, []string{"xian"})
	   	// fmt.Println(*p)

	   	// 结构体方法
	   	// functions()

	   	// 结构体继承
	   	// extends()

	   	//结构体与JSON序列化
	   	// structjson() */

	// 接口
	// interfacefunc()

}

/* //结构体方法
type student struct {
	name string
	age  int
}

func newStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}
func (s *student) setName(name string) {
	s.name = name
}
func (s student) eat() {
	fmt.Println(s.name + " 吃东西")
}

func functions() {
	s := newStudent("zx", 12)
	s.eat()
  s.setName("xm")
	s.eat()
} */

/* // 结构体继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Println(a.name + "移动")
}

type Dog struct {
	age     int
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s叫%v", d.name, d.age)
}
func extends() {
	d := &Dog{
		age: 12,
		Animal: &Animal{
			name: "ddd",
		},
	}
	d.move()
	d.wang()
}
*/

/* // 结构体与JSON序列化
type Student struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}
type classes struct {
	Title    string    `json:"title"`
	Students []Student `json:"students"`
}

func structjson() {
	c := &classes{
		Title:    "101",
		Students: make([]Student, 0, 50),
	}
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("%d", i),
			Age:  i + 10,
			Sex:  "man",
		}
		c.Students = append(c.Students, stu)
	}
	fmt.Println(c)
	data, _ := json.Marshal(*c)
	fmt.Println(string(data))
	// 反序列化
	str := `{"title":"101","students":[{"name":"0","sex":"man","age":10}]}`
	c1 := &classes{}
	json.Unmarshal([]byte(str), c1)
	fmt.Println((*c1).Title)
}
*/

/*
接口

	type active interface {
		focus()
		hover()
	}

	type pointer struct {
		x int
		y int
	}

	func (p *pointer) focus() {
		fmt.Printf("[%v,%v]focus激活\n", p.x, p.y)
	}

	func (p *pointer) hover() {
		fmt.Printf("[%v,%v]hover激活\n", p.x, p.y)
	}

	func focus(a active) {
		a.focus()
	}

	func interfacefunc() {
		p := &pointer{
			x: 1,
			y: 2,
		}
		focus(p)
	}
	// 		// 接口型函数只能应用于接口内部只定义了一个方法的情况
	// type Getter interface {
	// 	Get(key string) ([]byte, error)
	// }
	// type GetterFunc func(key string) ([]byte, error)

	// func (f GetterFunc) Get(key string) ([]byte, error) {
	// 	return f(key)
	// }
	// func GetFromSource(getter Getter, key string) []byte {
	// 	buf, err := getter.Get(key)
	// 	if err == nil {
	// 		return buf
	// 	}
	// 	return nil
	// }

	// func test(key string) ([]byte, error) {
	// 	return []byte(key), nil
	// }

	// func interfacefunc() {
	// 	// 将 test 强制类型转换为 GetterFunc，GetterFunc 实现了接口 Getter，是一个合法参数
	// 	// GetFromSource(GetterFunc(test), "hello")
	// 	// 实现了 Getter 接口的结构体作为参数
	// 	// GetFromSource(GetterFunc(func(key string) ([]byte, error) {
	// 	// 	return []byte(key), nil
	// 	// }), "hello")
	// 	// 既能够将普通的函数类型（需类型转换）作为参数，也可以将结构体作为参数，
	// 	// 使用更为灵活，可读性也更好，这就是接口型函数的价值
	// }
*/

/* func flagFunc() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	// //返回命令行参数后的其他参数
	// fmt.Println(flag.Args())
	// //返回命令行参数后的其他参数个数
	// fmt.Println(flag.NArg())
	// //返回使用的命令行参数个数
	// fmt.Println(flag.NFlag())
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			if index != 0 {
				s := strings.Split(arg, "=")
				fmt.Println(s[0], s[1])
			}
		}
	}

} */

/*
	 const (
			n1 = iota //0
			n2        //1
			n3        //2
			n4        //3
		)

const (

		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)

const (

		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)


rune 类型：代表一个Unicode码点，等同于int32，用于处理Unicode字符。
byte类型：作为uint8的别名，常用于处理字节数据。
s2 := "小白兔"
s3 := []rune(s2)        //把字符串强制转成rune切片
s3[0] = '大'             //注意 这里需要使用单引号的字符，而不是双引号的字符串
fmt.Println(string(s3)) //把rune类型的s3强转成字符串
*/
