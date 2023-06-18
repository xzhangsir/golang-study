package main

import (
	"fmt"
	"errors"
)

func intSum(x,y int) int{
	return x + y
}
func intSumSome(x ...int)(sum int){
	for _,v := range x{
		sum += v
	}
	return 
}
func calc(x,y int)(int,int){
	sum := x + y
	sub := x - y 
	return sum,sub
}
// func calc(x,y int)(sum,sub int){
// 	sum = x + y
// 	sub = x - y 
// 	return
// }

// 定义函数类型
// type calculation func(int, int) int

// 高阶函数 
// 函数作为参数
func add(x,y int)int{
	return x + y
}
func calc2(x,y int,opt func(int,int) int) int{
	return opt(x,y)
}
// 函数作为返回值
func do(s string)(func (int,int) int,error){
	switch s {
		case "+":
			return add,nil
			
		case "-":
			return func (x,y int) int{
				return x-y
			},nil
		default:
			err:= errors.New("无法识别")
			// err := "无法识别"
			return nil,err
	}
}

func main()  {
	// sum := intSum(1,2)
	// sum := intSumSome(2,3,4)
	// fmt.Println(sum)

	// sum,sub := calc(10,12)
	// fmt.Println(sum,sub)

	// ret := calc2(4,3,add)
	// fmt.Println(ret)

	// fmt.Println(do("+"))

	// fmt.Println(f1()) //5
	// fmt.Println(f2()) //6
	// fmt.Println(f3()) //5
	// fmt.Println(f4()) //5

	funcA()
	funcB()
	funcC()
}
// defer语句 可以理解为压栈
// func f1() int {
// 	x := 5
// 	defer func() {
// 		x++
// 	}()
// 	return x
// }
// func f2()(x int){
// 	defer func(){
// 		x++
// 	}()
// 	return 5
// }
// func f3()(y int){
// 	x := 5
// 	defer func(){
// 		x++
// 	}()
// 	return x
// }
// func f4()(x int){
// 	defer func(x int){
// 		x++
// 	}(x)
// 	return 5
// }

// 异常处理
// panic/recover
func funcA(){
	fmt.Println("func A")
}

func funcB(){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("revocer in B")
		}
	}()
	panic("panic in B")
}
func funcC(){
	fmt.Println("func C")
}


