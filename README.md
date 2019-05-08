# ananb
一个很简单的消息队列


## 快速入手

```go
package main

import (
    "github.com/Justyer/ananb"
    "fmt"
)

func main() {
    tq := ananb.NewQueue().SetCap(100).Init()

    tq.Push([]byte{'hi'})
    task, _ := tq.Pop()
    fmt.Println(string(task))
}
```

## 方法

```go
// 打入消息
Push([]byte) error

// 打入一堆消息
PushMany([][]byte) error

// 打入消息，如果队列满了会阻塞
MustPush([]byte)

// 打入一堆消息，如果队列满了会阻塞
MustPushMany([][]byte)

// 从队列中取一条消息, 如果数量不够则返回错误
Pop() ([]byte, error)

// 从队列中取一条消息, 如果mustGet为真时，即使数量不够也会取出，且不会报错，但剩余数量必须>0
PopMany(count int, mustGet bool) ([]byte, error)
```