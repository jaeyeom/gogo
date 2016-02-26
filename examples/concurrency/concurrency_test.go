package concurrency

import "fmt"

func Example_goroutine() {
	go func() {
		fmt.Println("In goroutine")
	}()
	fmt.Println("In main routine")
	// Non-deterministic!
}

func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 1
	// 2
	// 3
}

func Example_simpleChannelForLoop() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Example_simpleChannelReturnChannel() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func BabyNames(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func Example_babyNames() {
	for n := range BabyNames("성정명재경", "준호우훈진") {
		fmt.Print(n, ", ")
	}
	// Output:
	// 성준, 성호, 성우, 성훈, 성진, 정준, 정호, 정우, 정훈, 정진, 명준, 명호, 명우, 명훈, 명진, 재준, 재호, 재우, 재훈, 재진, 경준, 경호, 경우, 경훈, 경진,
}

func Example_closedChannel() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 0
	// 0
	// 0
}
