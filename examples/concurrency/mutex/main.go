package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
	wg  *sync.WaitGroup
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
	c.wg.Done()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

// an unsafe Counter, may cause a fetal error of concurrent map read/write
type UnSafeCounter struct {
	v  map[string]int
	wg *sync.WaitGroup
}

func (c *UnSafeCounter) Inc(key string) {
	c.v[key]++
	c.wg.Done()
}

func (c *UnSafeCounter) Value(key string) int {
	return c.v[key]
}

func main() {
	cf := flag.Int("count", 3, "inc count")
	flag.Parse()

	count := *cf
	fmt.Println("increment count: ", count)

	fmt.Println("SafeCounter demo:")
	counter := SafeCounter{v: make(map[string]int), wg: &sync.WaitGroup{}}
	counter.wg.Add(count)
	fmt.Printf("Inc() by %d start: %v\n", count, time.Now())
	for i := 0; i < count; i++ {
		go counter.Inc("somekey")
	}
	fmt.Printf("Inc() by %d finish: %v\n", count, time.Now())

	//time.Sleep(time.Second * 5)
	counter.wg.Wait()
	fmt.Println(counter.Value("somekey"))

	fmt.Println("============")
	fmt.Println("Counter demo:")
	counter2 := UnSafeCounter{v: make(map[string]int), wg: &sync.WaitGroup{}}
	counter2.wg.Add(count)
	fmt.Printf("Inc() by %d start: %v\n", count, time.Now())
	for i := 0; i < count; i++ {
		go counter2.Inc("somekey")
	}
	fmt.Printf("Inc() by %d finish: %v\n", count, time.Now())

	//time.Sleep(time.Second * 5)
	counter2.wg.Wait()
	fmt.Println(counter2.Value("somekey"))
}
