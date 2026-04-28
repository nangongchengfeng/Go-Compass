package main

import (
	"fmt"
	"sync"
	"time"
)

// 基础 Goroutine 示例
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 使用 WaitGroup 等待
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 函数结束时调用 Done
	fmt.Printf("Worker %d 开始工作\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d 完成工作\n", id)
}

// 使用 Channel 通信
func produce(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for num := range ch {
		fmt.Printf("消费: %d\n", num)
		time.Sleep(300 * time.Millisecond)
	}
}

// Select 示例
func server1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "来自 Server 1 的响应"
}

func server2(ch chan string) {
	time.Sleep(800 * time.Millisecond)
	ch <- "来自 Server 2 的响应"
}

// Mutex 互斥锁示例
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	fmt.Println("=== 基础 Goroutine ===")
	go sayHello("Alice")
	go sayHello("Bob")
	go sayHello("Charlie")

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	fmt.Println("=== WaitGroup 等待多个 Goroutine ===")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println()

	fmt.Println("=== Channel 生产者-消费者 ===")
	ch := make(chan int, 2) // 带缓冲的 Channel
	go produce(ch)
	go consume(ch)
	time.Sleep(3 * time.Second)
	fmt.Println()

	fmt.Println("=== Select 多路复用 ===")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)

	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("请求超时")
	}
	fmt.Println()

	fmt.Println("=== Mutex 互斥锁 ===")
	counter := &Counter{}
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter.Increment()
		}()
	}
	wg2.Wait()
	fmt.Printf("最终计数: %d\n", counter.Value())
}
