
### go ###

一、[环境安装](https://www.cnblogs.com/marshhu/p/11841042.html)

1.1 系统变量配置
   
>+GOROOT、PATH(bin)、GOPATH

1.2 创建工作空间

```textmate
Go代码必须放在工作空间中，实际上就是一个目录，其中包含三个子目录：
  src目录：包含Go的源码文件（每个目录都对应一个包）
  pkg目录：包含包对象
  bin目录：包含可执行对象
  GOPATH环境变量的设置：
GOPATH环境变量指向了工作空间的位置
```

1.3 [golang 自动下载所有依赖包](https://www.cnblogs.com/landv/p/10948227.html)

>*  go get -d -v ./...

1.4 [go mod 自动依赖方案](https://www.jb51.net/article/208264.htm)

>+ go env -w GOPROXY=https://goproxy.io,direct
>+ go env -w GO111MODULE=on
>注：可以用go env -u 恢复初始设置；GOPROXY的值应该还可以是https://mirrors.aliyun.com/goproxy/  或 https://goproxy.cn

```
go 自动安装项目依赖包
1,初始化
go mod init 项目名
2,直接安装
go get -d -v ./...
3, setting module: Enable Go modules integration 
```


###二、[beego](https://www.tizi365.com/archives/104.html) ###


