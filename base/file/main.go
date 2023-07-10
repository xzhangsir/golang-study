package main

// *******读取文件 基础版*******
// func main() {
// 	// 只读方式打开当前目录下的test.txt文件
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		fmt.Println("打开文件错误")
// 		return
// 	}
// 	defer file.Close()
// 	b := make([]byte, 128)
// 	// 使用Read方法读取数据 最多读取len(b)个字节
// 	n, err := file.Read(b)
// 	if err == io.EOF {
// 		fmt.Println("文件读取完了")
// 		return
// 	}
// 	if err != nil {
// 		fmt.Println("读取文件错误, err:", err)
// 		return
// 	}
// 	fmt.Printf("读取了%d字节数据\n", n)
// 	fmt.Println(string(b[:n]))
// }

// *******循环读取文件*******
// func main() {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		fmt.Println("打开文件错误")
// 		return
// 	}
// 	defer file.Close()
// 	var con []byte
// 	for {
// 		b := make([]byte, 2)
// 		n, err := file.Read(b)
// 		if err == io.EOF {
// 			fmt.Println("文件读取完成")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("读取文件错误, err:", err)
// 			return
// 		}
// 		con = append(con, b[:n]...)
// 	}
// 	fmt.Println(string(con))
// }

//******bufio读取文件******
// func main() {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		fmt.Println("打开文件错误")
// 		return
// 	}
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			if len(line) != 0 {
// 				fmt.Println(line)
// 			}
// 			fmt.Println("文件读完了")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("读取文件错误, err:", err)
// 			return
// 		}
// 		fmt.Print(line)
// 	}
// }

// ***** 读取整个文件 ****

// func main() {
// 	// ioutil.ReadFile在v1.16及更高版本中已废弃
// 	con, err := os.ReadFile("test.txt")
// 	if err != nil {
// 		fmt.Println("读取文件错误, err:", err)
// 		return
// 	}
// 	fmt.Println(string(con))
// }

// **** 文件写入操作 ****
// func OpenFile(name string, flag int, perm FileMode) (*File, error) {
// os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
// name：要打开的文件名
// flag：打开文件的模式
//       os.O_WRONLY	只写
//       os.O_CREATE	创建文件
//       os.O_RDONLY	只读
//       os.O_RDWR	读写
//       os.O_TRUNC	清空
//       os.O_APPEND	追加
// perm：文件权限，一个八进制数。rwx
// }

// func main() {
// 	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0644)
// 	if err != nil {
// 		fmt.Println("打开文件, err:", err)
// 		return
// 	}
// 	defer file.Close()
// 	// file.WriteString("xixixihhhh")//直接写入字符串数据

// 	// str := "为人民服务"
// 	// file.Write([]byte(str)) //写入字节切片数据
// }

// func main() {
// 	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0644)
// 	if err != nil {
// 		fmt.Println("打开文件, err:", err)
// 		return
// 	}
// 	defer file.Close()
// 	writer := bufio.NewWriter(file)
// 	for i := 0; i < 10; i++ {
// 		writer.WriteString("hello\n") //将数据先写入缓存
// 	}
// 	writer.Flush() //将缓存中的内容写入文件
// }

// 简化了打开、写入和关闭文件的过程
// func main() {
// 	str := "hello"
// 	err := ioutil.WriteFile("test.txt", []byte(str), 0666)
// 	if err != nil {
// 		fmt.Println("写入文件失败, err:", err)
// 		return
// 	}
// }

// **** 文件copy****
// func main() {
// 	srcFile := "test.txt"
// 	dstFile := "copytest.txt"
// 	con, err := os.ReadFile(srcFile)
// 	if err != nil {
// 		fmt.Println("读取文件错误, err:", err)
// 		return
// 	}
// 	err = ioutil.WriteFile(dstFile, con, 0666)
// 	if err != nil {
// 		fmt.Println("写入文件失败, err:", err)
// 		return
// 	}
// 	fmt.Println("文件复制成功")
// }
