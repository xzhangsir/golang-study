package main

import (
	"fmt"
	// "strings"
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
}


