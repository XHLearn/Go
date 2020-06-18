# 三、编辑器、集成开发环境
## 3.1 vscode智能提示
1. 自动完成列表(使用gocode):
    - `go get -u -v github.com/nsf/gocode`
2. 快速提示函数定义、跳转定义(使用godef):
    - `go get -u -v github.com/rogpeppe/godef`
3. 查找引用(使用go-find-references)：
    - `go get -u -v github.com/lukehoban/go-find-references`
4. 代码大纲(使用go-outline)：
    - `go get -u -v github.com/lukehoban/go-outline`
5. 重命名(使用gorename)
    - `go get -u -v golang.org/x/tools/cmd/gorename`
6. 调试(使用delve)
    - `go get -u -v github.com/derekparker/delve`
7. 安装golint:
    - `go get -u -v github.com/golang/lint/golint`
8. 代码格式化（使用goreturns或goimports或gofmt）
    - `go get -u -v sourcegraph.com/sqs/goreturns`

## 3.5 格式化代码
- gofmt –w program.go：
- gofmt -w *.go ：会格式化并重写所有 Go 源文件;
- gofmt map1 ：会格式化并重写 map1 目录及其子目录下的所有 Go 源文件

## 3.6 生产代码文档
- go doc package 获取包的文档注释
- go doc package/subpackage   获取子包的文档注释
- go doc package function   获取某个函数在某个包中的文档注释

## 3.7 go程序执行顺序
1. 按顺序导入所有被 main 包引用的其它包,然后在每个包中执行如下流程:
2. 如果该包又导入了其它的包,则从第一步开始递归执行,但是每个包只会被导入一次。
3. 然后以相反的顺序在每个包中初始化常量和变量,如果该包含有 init 函数的话,则调用该函数。
4. 在完成这一切之后,main 也执行同样的过程,最后调用 main 函数开始执行程序。

## 4.5 格式化说明符
- fmt.Printf()

> %d 用于格式化整数
> %x和%X 用于格式化 16 进制表示的数字
> %g 用于格式化浮点型
> %f 输出浮点数
> %e 输出科学计数表示法
> %0d 用于规定输出定长的整数,其中开头的数字 0 是必须的。
> %n.mg 用于表示数字 n 并精确到小数点后 m 位,除了使用 g 之外,还可以使用 e 或者 f
> %p 指针的格式化标识符
> %T 输出类型

## 6.5 传递变长参数
- `func myFunc(a, b, arg ...int) {}` 其中arg为[]int类型
- `myFunc(a, b, arg...)` 调用也可以这样调用
- `myFunc(a, b, arg1, arg2, arg3)`

## 6.4 defer 和追踪
> 关键字 defer 允许我们推迟到函数返回之前(或任意位置执行return语句之后)一刻才执行某个语句或函数
> 为什么要在返回之后才执行这些语句,因为return语句同样可以包含一些操作,而不是单纯地返回某个值
> 当有多个 defer 行为被注册时,它们会以注册逆序执行

## 7.2 make创建切片
- `var slice []type = mak([]type, len, cap)`
- len:是数组的长度并且也是slice的初始长度
- `make([]int, 50, 100)` 等同于 `new([100]int)[0:50]`
- new() 和 make() 的区别:
    - new(T) 为每个新的类型T分配一片内存,初始化为 0 并且返回类型为*T的内存地址:这种方法 返回一个指向类型为
T,值为 0 的地址的指针,它适用于值类型如数组和结构体(参见第 10 章);它相当于
&T{}
。
    - make(T) 返回一个类型为 T 的初始值,它只适用于3种内建的引用类型:切片、map 和 channel(参见第 8 章,第 13
章)。

## 7.4
- go字符串是不可改变的，例如:str[index]='x',会报错；
- 正确做法：
```go
s := "hello"
c := []byte(s)
c[0] = 'c'
s2 := string(c) // s2 == "cello"
```

