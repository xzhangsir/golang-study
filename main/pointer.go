package main

import (
	"fmt"
)

// &取地址 *根据地址取值
func main(){

	// a := 12
	// b := &a 
	// fmt.Println(&a,b,*b,a) //地址 地址 12 12
	// *b = 22
	// fmt.Println(&a,b,*b,a) //地址 地址 22 22
	

	var a *int
	a = new(int) //对a进行初始化后 才有内存空间 才可以赋值
	*a = 100
	fmt.Println(a)

}