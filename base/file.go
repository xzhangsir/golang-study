package base

/* // 循环读取整个文件
func FileFunc() {
	file, err := os.Open("./temp/index.txt")
	if err != nil {
		fmt.Println("打开文件错误", err)
		return
	}
	defer file.Close()
	var con []byte
	b := make([]byte, 4)
	for {
		n, err := file.Read(b)
		if err == io.EOF {
			fmt.Println("文件读取完了")
			break
		}
		if err != nil {
			fmt.Println("读取文件错误, err:", err)
			return
		}
		con = append(con, b[:n]...)
		// fmt.Printf("读取了%d字节数据\n", n)
		// fmt.Println(string(b[:n]))
	}
	fmt.Println(string(con))
} */

/* // ReadFile函数 读取完整的文件
func FileFunc() {
	content, err := os.ReadFile("./temp/index.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
} */

/* // 文件写入操作
// func OpenFile(name string, flag int, perm FileMode) (*File, error) {
// os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
// name：要打开的文件名
// flag：打开文件的模式
//
//	os.O_WRONLY	只写
//	os.O_CREATE	创建文件
//	os.O_RDONLY	只读
//	os.O_RDWR	读写
//	os.O_TRUNC	清空
//	os.O_APPEND	追加
//
// perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01
// }
func FileFunc() {
	file, err := os.OpenFile("./temp/index.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("打开文件, err:", err)
		return
	}
	defer file.Close()
	file.WriteString("为人民服务") //直接写入字符串数据
	str := "weirenmingfuwu"
	file.Write([]byte(str)) //写入字节切片数据
} */

/* // bufio读取文
func FileFunc() {
	file, err := os.Open("./temp/index.txt")
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("读取文件错误, err:", err)
			return
		}
		fmt.Print(line)
	}
}
*/

/* // bufio 写入文件
func FileFunc() {
	file, err := os.OpenFile("./temp/index.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("打开文件, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}
*/

/* // 简化了打开、写入和关闭文件的过程
func FileFunc() {
	str := "hello"
	err := ioutil.WriteFile("./temp/index.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("写入文件失败, err:", err)
		return
	}
}
*/

/* // 实现文件copy
func FileFunc() {
	srcFile := "test.txt"
	dstFile := "copytest.txt"
	con, err := os.ReadFile(srcFile)
	if err != nil {
		fmt.Println("读取文件错误, err:", err)
		return
	}
	err = ioutil.WriteFile(dstFile, con, 0666)
	if err != nil {
		fmt.Println("写入文件失败, err:", err)
		return
	}
	fmt.Println("文件复制成功")
}
*/
