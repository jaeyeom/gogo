package concurrency

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/net/context"
)

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	ctx := context.Background()
	for num := range PlusOne(ctx, PlusOne(ctx, c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

func Example_contextSwitching() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	// Non-deterministic!
}

func ExampleFanIn3() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)
	for n := range FanIn3(c1, c2, c3) {
		fmt.Print(n, ",")
	}
	// Non-deterministic!
}

func ExamplePlusOne_consumeAll() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	ctx := context.Background()
	nums := PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}
	time.Sleep(100 * time.Millisecond)
	// fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	log.Println("NumGoroutine: ", runtime.NumGoroutine())
	for _ = range nums {
		// Consume all nums
	}
	time.Sleep(100 * time.Millisecond)
	// fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	log.Println("NumGoroutine: ", runtime.NumGoroutine())
	// Output:
	// 8
	// 18
}

func ExamplePlusOne_withCancel() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	nums := PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			cancel()
			break
		}
	}
	// Output:
	// 8
	// 18
}
