[TOC]

### Golang与C语言进行交互

工具cgo提供了对FFI(外部函数接口)的支持，能够使用Golang代码安全地调用C语言库，你可以访问：[CGO文档主页](http://golang.org/cmd/cgo)。cgo会替代Golang编译器来产生可以组合在同一个包的Golang和C代码。在实际开发中一般使用cgo创建单独的C代码包。

如果你想要在你的程序中使用cgo，则必须在单独的一行使用`import "C"`来导入，一般来说你可能还需要`import "unsafe"`。

然后，你可以在`import "C"`之前使用注释(单行或多行注释均可)的形式导入C语言库(甚至有效的C语言代码)，它们之间没有空行，例如：

```go
// #include <stdio.h>
// #include <stdlib.h>
import "C"
```

名称“C”并不属于标准库的一部分，这只是cgo集成的一个特殊名称用于引用C的命名空间。在这个命名空间里所包含的C类型都可以被使用，例如`C.uint`、`C.long`等等，还有libc中的函数`C.random()`等也可以被调用。

当你想要使用某个类型作为C中函数的参数时，必须将其转换为C语言中的类型，反之亦然，例如：

```go
var i int
C.uint(i) 		// 从 Go 中的 int 转换为 C 中的无符号 int
int(C.random()) // 从 C 中 random() 函数返回的 long 转换为 Go 中的 int
```

下面的2个Golang函数`Random()`和`Send()`分别调用了C中`C.random()`和`C.srandom()`。

```go
package rand

// #include <stdlib.h>
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
```

C当中并没有明确的字符串类型，如果你想要将一个`string`类型的变量从Golang转换到C时，可以使用`C.CString(s)`；同样，可以使用`C.GoString(cs)`从C转换到Golang中的`string`类型。

Golang的内存管理机制无法管理通过C代码分配的内存。

开发人员需要通过手动调用`C.free`来释放变量的内存：

```go
defer C.free(unsafe.Pointer(Cvariable))
```

这一行最好紧跟在使用C代码创建某个变量之后，这样就不会忘记释放内存了。下面的代码展示了如何使用cgo创建变量、使用并释放其内存：

```go
package print

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}
```

**构建cgo包**

你可以在使用Makefile文件(因为我们使用了一个独立的包)，除了使用变量`GOFILES`之外，还需要使用变量`CGOFILES`来列出需要使用cgo编译的文件列表。例如，下面的代码就可以使用包含以下内容的`Makefile`文件来编译，你可以使用`gomake`或`make`:

```shell
include $(GOROOT)/src/Make.inc
TARG=rand
CGOFILES=\
c1.go\
include $(GOROOT)/src/Make.pkg
```

### Golang与C++进行交互

SWIG（简化封装器和接口生成器）支持在Linux系统下使用Golang代码调用C或者C++代码。这里有一些使用SWIG的注意事项：

* 编写需要封装的库的SWIG接口；
* SWIG会产生C的存根函数；
* 这些库可以使用cgo来调用；
* 相关的Golang文件也可以被自动生成

这类接口支持方法重载、多重继承已经使用Golang代码实现C++的抽象类。

目前使用SWIG存在的一个问题是它无法支持所有的C++库，比如TObject.h。