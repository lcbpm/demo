### sync.WaitGroup

```text
var wg sync.WaitGroup 用于创建一个 WaitGroup 实例，它在 Go 并发编程中非常有用，尤其是在处理多个 Goroutine 的同步时。
主要作用：
等待一组 Goroutine 完成：WaitGroup 可以用来等待多个 Goroutine 执行完成，通常用于确保在主 Goroutine 结束前所有工作 Goroutine 都已完成。
使用步骤：
添加计数：使用 wg.Add(n) 增加等待计数 n。
Goroutine 内部调用 wg.Done()：在每个 Goroutine 中执行完任务后调用 wg.Done()，减少计数。
等待所有 Goroutine 完成：使用 wg.Wait() 在主 Goroutine 中等待计数器归零。
```

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers completed")
}

```