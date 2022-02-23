
### go ###

一、[环境安装](https://www.cnblogs.com/marshhu/p/11841042.html)

1.1 系统变量配置
   
>+GOROOT(安装目录)、PATH(bin)、GOPATH(工作目录)

1.2 创建工作空间

```textmate
Go代码必须放在工作空间中，实际上就是一个目录，其中(GOPATH)包含三个子目录：
  src目录：包含Go的源码文件（每个目录都对应一个包）
  pkg目录：包含包对象
  bin目录：包含可执行对象
  GOPATH环境变量的设置：
GOPATH环境变量指向了工作空间的位置
```

1.3 [golang 自动下载所有依赖包](https://www.cnblogs.com/landv/p/10948227.html)

>*  go get -d -v ./...

```text

go get -d -v ./...
-d标志只下载代码包，不执行安装命令；
-v打印详细日志和调试日志。这里加上这个标志会把每个下载的包都打印出来；
./...这个表示路径，代表当前目录下所有的文件。
补充：goland自动下载所有依赖

项目中使用了go.mod时可以使用以下命令自动下载全部依赖
方法一
1
go get -d -v ./...
方法二
1
go mod tidy
```

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


### 三、[GF 脚手架](https://www.cnblogs.com/haima/p/13041257.html) ###
```textmate
gf gen model ./app/model -c config/config.toml -p t_ -t t_encrypt_test

命令说明

./app/model #在model生成的路径
-c config/config.toml #在这个配置里找database数据库连接配置 需要写好mysql的配置信息

```


### 学习指南 ###
[go语言四十二章经](https://www.jianshu.com/nb/29056963)
[|| github](https://github.com/ffhelicopter/Go42)
[go语言高级编程 + go语言核心编程](https://studygolang.com/articles/26505/comment/38656)
[|| pan](下载: https://pan.baidu.com/s/1LMpyblrNnp9dyG7zUh9--A  提取码: m9qb)


### 知识点 ###
>+ 位运算补码
>+ a, b, c := 5, 7, "abc" 简式声明，右边以相同顺序复制给左边变量
>+ _, b = 5, 7， _空白标识符用于抛弃值，Go语言有个强制规定，在函数内一定要使用声明的变量， 引入包也需要使用
>+ Go中的数组是数值，因此当你向函数中传递数组时，函数会得到原始数组数据的一份复制。如果你打算更新数组的数据，可以考虑使用数组指针类型。
> ? 函数的写法 func(xx) {}(&x), 被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间
> nil是派生类的零值， 不指定变量的类型编译器无法编译你的代码；  var x = nil // 错误
> nil的slice可以添加元素，对map会产生panic,   var m map[string]int ;  m["one"] = 1 //error
>+ 常量格式：const identifier [type] = value，  type可以省略，通过编译器进行推断，常量智能是布尔、数字（整数型、浮点型和复数）、字符串, iota是一个预定义常量，可被编译器修改，首次重置为0，下一个const出现前+1
>+ 局部、全局、简式变量（使用 := 定义的变量）会遮住全局变量重新定义该变量即无法修改; if是由隐式代码块和显示代码逻辑组成， 代码块的作用域限定在块内；闭包函数可以理解为一个代码块，并且他可使用包含它的函数内的变量；
>+ 函数保留给外部包，需要符合大驼峰命名法
>+ go也可以使用分好作为语句的结束，但一般会省略分号。 
>+ [rune类型的作用](https://www.jianshu.com/p/4fbf529926ca)
>+ 特殊import: 1、import( . "fmt" )  函数调用不需要包名;  2、import( f "fmt" ) 别名调用  3、 _ 操作是引入某个包，但不直接使用包里的函数，而是调用该包里面的init函数，比如下面的mysql包的导入
>+ go install会安装到$GoPATH/pkg/目录下 ， go get -u 和前者的区别
>+ init()函数在main调用前进行初始化，无法被其他函数调用 
>+ 注释doc生成： godoc -http=:6060 -goroot="." 
>+ 1.11 新增了对模块的支持，希望借此解决“包依赖管理”go mod管理，  GOPATH是无意义的，不过它还是会把下载的依赖储存在 GOPATH/pkg/mod 中。 
>+ string字符串长度问题， Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。 字符串拼接 strings.Join([]string{"hello", "world"}, ", ")， bytes.Buffer, strings.Builder
>+ [string类型](https://www.jianshu.com/p/19b2b20a7434) 常用标准库：bytes、strings、strconv和unicode包 
>+ [数组类型](https://www.jianshu.com/p/2a8c7c8cd82a) go语言中数组是一种值类型，不是数组指针，避免传递大数组给函数，值传递会消耗很多内存， 可采用数组的指针或切片； 数组大小最大为 2Gb
>+ [slice类型](https://www.jianshu.com/p/d9d9766a8a5a) slice是一个引用类型，是一个可变长度的数组； 使用内置函数make()可以给切片初始化；  var identifier []type（不需要说明长度）
> 创建方式 var x = []int{2, 3, 5, 7, 11}， 2、 v := make([]int, 10, 50)， 数组生成切片a := [5]int{1, 2, 3, 4, 5} --> t := a[1:3:5], 索引不能超过长度
> 切片重组，改变切片长度得到新的切片  copy(res, raw[:3])，  append()内置函数 追加返回新的切片s0 := []int{0, 0}，s1 := append(s0, 2)； append()函数操作后，有没有生成新的slice需要看原有slice的容量是否足够。
> [字典](https://www.jianshu.com/p/15f9d6f5935d) map是引用类型var map1 map[keytype]valuetype， 通过使用空接口类型，我们可以存储任意值； 初始化var map1 = make(map[keytype]valuetype)
> range语句中生成的数据是真实集合元素的拷贝，不是原有元素的引用。
>+ [方法](https://www.jianshu.com/p/f960db2469a7)
>+ [同步与锁](https://www.jianshu.com/p/90f7626d83e2) 
>+ [net/http](https://www.jianshu.com/p/fedd86adc68b) 
>+ [常用函数](https://vimsky.com/examples/usage/reflect-complex-function-in-golang-with-examples.html) 