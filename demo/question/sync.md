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

### sync.Map

```text
并发安全的 map，提供高效的读写操作，避免了自己实现锁机制。
```

```go
var m sync.Map

func main() {
    m.Store("key", "value")

    value, ok := m.Load("key")
    if ok {
        fmt.Println(value)
    }

    m.Range(func(k, v interface{}) bool {
        fmt.Println(k, v)
        return true
    })
}

```

### sync.RWMutex

```text
读写互斥锁，允许多个读取锁定，但写入锁定是互斥的。
```

```go
var rw sync.RWMutex

func read() {
    rw.RLock()
    fmt.Println("Reading")
    time.Sleep(1 * time.Second)
    rw.RUnlock()
}

func write() {
    rw.Lock()
    fmt.Println("Writing")
    time.Sleep(1 * time.Second)
    rw.Unlock()
}

func main() {
    for i := 0; i < 3; i++ {
        go read()
    }

    for i := 0; i < 3; i++ {
        go write()
    }

    time.Sleep(3 * time.Second)
}

```

### sync.Cond

```text
条件变量，用于 goroutine 间的复杂同步，可以使 goroutine 等待某个条件满足。
```

```go
var mu sync.Mutex
var cond = sync.NewCond(&mu)

func main() {
    for i := 0; i < 10; i++ {
        go func(i int) {
            mu.Lock()
            cond.Wait()
            fmt.Println("Goroutine", i)
            mu.Unlock()
        }(i)
    }

    time.Sleep(time.Second)
    fmt.Println("Broadcasting...")
    cond.Broadcast()
    time.Sleep(time.Second)
}

```

### sync.Once

```text
确保某些操作只执行一次。通常用于单例模式或初始化操作。
```

```go
var once sync.Once

func initialize() {
    fmt.Println("Initialized")
}

func main() {
    for i := 0; i < 10; i++ {
        go once.Do(initialize)
    }
    time.Sleep(time.Second)
}

```