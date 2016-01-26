package concurrency

import (
	"sync"

	"golang.org/x/net/context"
)

// PlusOne returns a channel of num + 1 for nums received from in.
func PlusOne(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

type IntPipe func(context.Context, <-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(ctx, c)
		}
		return c
	}
}

var PlusTwo = Chain(PlusOne, PlusOne)

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Distribute(p IntPipe, n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(ctx, in)
		}
		return FanIn(cs...)
	}
}

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCnt--
		return openCnt == 0
	}
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			}
		}
	}()
	return out

}
