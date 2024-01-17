尽管Golang编译器产生的是本地可执行代码，这些代码仍旧运行在Golang的runtime（这部分的代码可以在`runtime`
包中找到）当中。这个runtime类似Java和.NET语言所用到的虚拟机，它负责管理包括内存分配、垃圾回收、栈处理、Goroutine、channel、切片、map和反射等等。

runtime主要由C语言编写，并且是每个Golang包的最顶级包。你可以在目录`$GOROOT/src/runtime`中找到相关内容。

**垃圾回收器**：Golang拥有简单却高效的标记——清除回收器。它的主要思想来源于IBM的可复用垃圾回收器，旨在