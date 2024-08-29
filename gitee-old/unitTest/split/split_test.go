package split

import (
	"testing"	
	"reflect"
)

// 直接在当前包下  go test  运行  go test -v  详细信息
// go  test -cover  测试代码覆盖率

// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
func TestSplit(t *testing.T){
	// got := Split("a::b::c", "::")         // 程序输入的结果
	// want := []string{"a", "b", "c"}    // 期望的结果
	// // 比较 want, got  是否相等
	// if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
	// 	t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	// }
  

  	// 测试组

	/*type test struct{
		input string
		sep string
		want []string
	}

	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"","河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}*/

	// 子测试
	// go test -run=TestSplit=2  单独跑某一个测试

	type test struct{
		input string
		sep string
		want []string
	}

	tests := map[string]test{
	"1":	{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
	"2":	{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
	"3":	{input: "abcd", sep: "bc", want: []string{"a", "d"}},
	"4":	{input: "沙河有沙又有河", sep: "沙", want: []string{"","河有", "又有河"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}


// 基准测试就是在一定的工作负载之下检测程序性能
// go test -bench=Split  执行
// go test -bench=Split -benchmem  查看执行内存的情况
func BenchmarkSplit(b *testing.B){
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}