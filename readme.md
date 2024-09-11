```
go mod init 项目名
go build // go run main.go
```

<!-- 依赖包下载 -->

```
go get -u github.com/q1mi/hello

or

先编辑go.mod文件 ，将依赖包和版本信息写入该文件
require github.com/q1mi/hello latest

然后执行  go mod download
```

<!-- 清除所有本地已缓存的依赖包数据 -->

```
 go clean -modcache
```

```
go work init ./xx ./xx
```

<!-- chat即时聊天系统  -->

```
 //在当前文件夹下
 go run main.go user.go
//在当前文件夹下
 go run client.go
```
