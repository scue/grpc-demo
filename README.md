# GRPC示例

这是一个练手的程序，从客户端发送命令给服务端进行调用，然后返回结果。

![](https://ws1.sinaimg.cn/large/6e22ca27gy1fruqksyeyvj20jo03st91.jpg)

# 运行方式

启动服务端:

```sh
go run server/main.go 
```

启动客户端:

```sh
go run client/main.go
```

# 编译`.proto`

我这边已经编译好了，不过如果你改动了，可以这样子处理：

```sh
cd backdog
protoc -I . --go_out=plugins=grpc:. backdog.proto
```

> Mac下的`protoc`安装方式：`brew install protobuf`

然后就会重新生成一个`backdog.pb.go`文件。