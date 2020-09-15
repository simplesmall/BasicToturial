# Go语言基础学习记录

简单记录Go语言的入门尝试及学习路径

### Get start !

```bas
下载GoLand -> 下载SDK ->导入SDK -> Create new project
```

### 项目结构 : 

```jso
BasicToturial:
├─.idea
├─Channel          协程测试
├─Metacognition    元认知[包括基础数据类型,控制语句等]
└─ModulesManTools  Modules管理工具设置和使用目录
```

课程原本是Go并发编程,目前更新至Jenkins单元测试部分.

```json
BasicToturial\Metacognition

Mode                LastWriteTime         Length Name
----                -------------         ------ ----
-a----        2020/9/15     17:05            821 BasicReview.go
-a----        2020/9/15     13:46            317 DeferTest.go
-a----        2020/9/15     12:26            623 ForTest.go
-a----        2020/9/15     11:27             72 HelloWorld.go
-a----        2020/9/15     13:40            317 IfTest.go
-a----        2020/9/15     12:28            543 MapTest.go
-a----        2020/9/15     11:29            655 Simple.go
-a----        2020/9/15     12:06            553 SliceTest.go
-a----        2020/9/15     11:56            575 Struct.go
-a----        2020/9/15     13:42            459 SwitchTest.go
```



```jso
BasicToturial\Channel

Mode                LastWriteTime         Length Name
----                -------------         ------ ----
-a----        2020/9/15     14:03            314 Basic.go
-a----        2020/9/15     14:15           1232 ContextTest.go
-a----        2020/9/15     14:09            565 PVTest.go
-a----        2020/9/15     14:14            701 SelectTest.go
```

### 代码预览 :

```go
// 代码示例  BasicToturial\Metacognition\BasicReview.go
package main

import "fmt"

func main() {
	fmt.Println("Testing main pak and main method")
	PointTest()
}

// 方法名首字母大写包外可访问
func Dev(source, target int) (handleSource, handleTarget int) {
	handleSource = source + 2
	handleTarget = 18989
	return
}

// 变量声明+初始化
var tempType int = 100

// 可省略类型自动推导
var temp = 100

// 短变量声明赋值,无法在函数外使用   name:= expression

func TestShortVar() {
	// 匿名函数接收error,后续不会用到
	value, _ := fmt.Println()
	// 声明的变量必须都使用到,否则编译报错
	fmt.Println(value)
}

//指针声明 var name *T
func PointTest() {
	name := "simple"
	// get pointer value
	p := &name
	fmt.Println("Name is : ", *p)
	// set pointer value
	*p = "Jack zhou"
	fmt.Println("Name is : ", *p)
}
```

