package main

import (
	"fmt"
    "os"
    "bufio"
    "io"
    "io/ioutil"
)
func main(){
  // readFromFile() //自定义一次读多少字节
	// readFileBufio() //一次读取一行
	// readFileIoutil() //一次读取所有
	// 文件写入
	// writeFile()
	// writeFileBufio()
	// writeFileIoutil()
	copy("./main.go","./test.txt")
}
func readFromFile(){
	// 打开文件
  fileObj,err := os.Open("./main.go")
  if err != nil{
  	fmt.Printf("打开文件错误,%v" , err)
  	return
  }
  // 关闭文件
  defer fileObj.Close()

  // var long = make([]byte,128)
  var long [128]byte //指定读取的长度
  for{
	  n,err := fileObj.Read(long[:])
	  if err != nil{
	  	fmt.Printf("读取文件错误,%v" , err)
	  	return
	  }
	  fmt.Printf("读取了%v个字节",n)
	  fmt.Printf("%v",string(long[:n]))
	  if n < 128 {
	  	fmt.Print("读取完成")
	  	return
	  }
  }
}
func readFileBufio(){
 	fileObj,err := os.Open("./main.go")
 	if err != nil{
 		fmt.Printf("打开文件错误%v",err)
 		return
 	}
 	defer fileObj.Close()
 	reder := bufio.NewReader(fileObj)
 	for{
		line,err := reder.ReadString('\n')
	 	if err == io.EOF{
	 		return
	 	}
	 	if err != nil{
	 		fmt.Printf("读取错误%v",err)
	 		return
	 	}
	 	fmt.Print(line)
 	}

 }
func readFileIoutil(){
	ret,err := ioutil.ReadFile("./main.go")
	if err != nil{
		fmt.Printf("读取错误%v",err)
		return
	}
	fmt.Print(string(ret))
}
func writeFile(){
	// os.O_WRONLY|os.O_CREATE|os.O_APPEND 创建文件并追加写入
	// os.O_WRONLY|os.O_CREATE|os.O_TRUNC  创建文件并清空写入
	fileObj,err := os.OpenFile("./test.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0600)
	if err != nil{
		fmt.Print("写入文件错误")
		return
	}
	defer fileObj.Close()
	fileObj.Write([]byte("写入字节切片数据\n"))
	fileObj.WriteString("直接写入字符串数据")
}
func writeFileBufio(){
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush()
}
func writeFileIoutil(){
	str := "hello 沙河"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
// 实现copy文件
func copy(target,newPath string){
	ret,err := ioutil.ReadFile(target)
	if err != nil{
		fmt.Printf("读取错误%v",err)
		return
	}
	err = ioutil.WriteFile(newPath, []byte(string(ret)), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}