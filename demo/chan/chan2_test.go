package _chan

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
}

// 生产者
func producer(jobs chan<- Job, quit <-chan bool) {
	id := 1
	for {
		select {
		case <-quit:
			close(jobs)
			fmt.Println("Producer stopped")
			return
		default:
			jobs <- Job{ID: id, Data: fmt.Sprintf("data-%d", id)}
			id++
			//time.Sleep(500 * time.Millisecond) // 模拟生产时间
		}
	}
}

// 消费者
func consumer(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Consumer %d processing job %d\n", id, job.ID)
		//time.Sleep(1 * time.Second) // 模拟处理时间
		results <- Result{JobID: job.ID, Output: fmt.Sprintf("processed data-%d", job.ID)}
	}
}

func Test(t *testing.T) {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	quit := make(chan bool)

	// 用于等待所有消费者完成工作
	var wg sync.WaitGroup

	// 启动生产者
	go producer(jobs, quit)

	// 启动多个消费者
	numConsumers := 3
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, jobs, results, &wg)
	}

	// 捕获终止信号以优雅退出
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		close(quit)
	}()

	// 等待所有消费者完成工作
	go func() {
		wg.Wait()
		close(results)
	}()

	// 输出结果
	for result := range results {
		fmt.Printf("Result: JobID %d, Output %s\n", result.JobID, result.Output)
	}

	fmt.Println("All jobs processed")
}
