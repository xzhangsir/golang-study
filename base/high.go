package base

// 一些进阶的知识
/* func High() {

	strMapFunc := func(arr []string, fn func(s string) string) []string {
		newArr := []string{}
		for _, it := range arr {
			newArr = append(newArr, fn(it))
		}
		return newArr
	}
	var s = []string{"xxx", "yyy"}
	x := strMapFunc(s, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(x)

	intReduceFunc := func(arr []int, fn func(sum, next int) int) int {
		sum := 0
		for _, v := range arr {
			sum = fn(sum, v)
		}
		return sum
	}
	var i = []int{1, 3, 5, 7, 9}
	sum := intReduceFunc(i, func(sum, next int) int {
		return sum + next
	})
	fmt.Println(sum)

	intFilterFunc := func(arr []int, fn func(i int) bool) []int {
		s := []int{}
		for _, v := range arr {
			if fn(v) {
				s = append(s, v)
			}
		}
		return s
	}
	f := intFilterFunc(i, func(i int) bool {
		return i > 5
	})
	fmt.Println(f)
}
*/

// 泛型实现 map/reduce/filter
/* func High() {

	type person struct {
		name  string
		age   int
		other string
	}
	list := []person{
		{name: "zx", age: 12},
		{name: "xm", age: 13},
	}

	others := mapFunc(list, func(v person) person {
		v.other = fmt.Sprintf("%s-%d", v.name, v.age)
		return v
	})
	fmt.Println(others)

	sumAge := reduceFunc(list, 0, func(init int, item person) int {
		return init + item.age
	})
	fmt.Println(sumAge)

	ages := filterFunc(list, true, func(item person) bool {
		return item.age > 12
	})
	fmt.Println(ages)

}

func mapFunc[T1 any, T2 any](arr []T1, f func(v T1) T2) []T2 {
	s := make([]T2, 0, len(arr))
	for _, v := range arr {
		s = append(s, f(v))
	}
	return s
}

func reduceFunc[T1 any, T2 any](arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}

func filterFunc[T any](arr []T, in bool, f func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, elem := range arr {
		choose := f(elem)
		if (in && choose) || (!in && !choose) {
			result = append(result, elem)
		}
	}
	return result
}
*/

func High() {
	/* // 随机字符串
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := func(n int) string {
		b := make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(b)
	}
	str := randomString(6)
	fmt.Println(str) */

	/* 	var builder strings.Builder    // 声明一个 Builder 变量
	   	builder.WriteString("Hello, ") // 追加字符串
	   	builder.WriteString("world!")  // 追加字符串
	   	fmt.Println(builder.String())  // 输出结果 */

	/* 	// 单例模式
	   	type singleton struct{}
	   	var instance *singleton
	   	var once sync.Once
	   	GetInstance := func() *singleton {
	   		// once.Do 并发的场景下只会执行一次
	   		once.Do(func() {
	   			instance = &singleton{}
	   		})
	   		return instance
	   	}
	   	GetInstance() */
}
