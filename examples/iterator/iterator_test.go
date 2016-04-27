package iterator

import (
	"testing"

	"golang.org/x/net/context"
)

const (
	size     = 10000
	eachTask = 1000
)

func doSomething(num int) int {
	for i := 0; i < eachTask; i++ {
		_ = i
	}
	return num
}

func BenchmarkCallback(b *testing.B) {
	// If f returns false, stop the iteration.
	iter := func(f func(num int) bool) {
		for i := 0; i < size; i++ {
			if !f(doSomething(i)) {
				break
			}
		}
	}

	for i := 0; i < b.N; i++ {
		iter(func(num int) bool {
			_ = num
			return true
		})
	}
}

func BenchmarkChannel(b *testing.B) {
	iter := func() chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 0; i < size; i++ {
				out <- doSomething(i)
			}
		}()
		return out
	}

	for i := 0; i < b.N; i++ {
		for num := range iter() {
			_ = num
		}
	}
}

func BenchmarkBufferedChannel(b *testing.B) {
	iter := func() chan int {
		out := make(chan int, 10)
		go func() {
			defer close(out)
			for i := 0; i < size; i++ {
				out <- doSomething(i)
			}
		}()
		return out
	}

	for i := 0; i < b.N; i++ {
		for num := range iter() {
			_ = num
		}
	}
}
func BenchmarkChannelWithContext(b *testing.B) {
	iter := func(ctx context.Context) chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 0; i < size; i++ {
				select {
				case out <- doSomething(i):
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < b.N; i++ {
		for num := range iter(ctx) {
			_ = num
			// cancel()
		}
	}
	_ = cancel // No-op
}

func BenchmarkBufferedChannelWithContext(b *testing.B) {
	iter := func(ctx context.Context) chan int {
		out := make(chan int, 10)
		go func() {
			defer close(out)
			for i := 0; i < size; i++ {
				select {
				case out <- doSomething(i):
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < b.N; i++ {
		for num := range iter(ctx) {
			_ = num
			// cancel()
		}
	}
	_ = cancel // No-op
}

func BenchmarkFunc(b *testing.B) {
	iter := func() func() (int, bool) {
		num, next := 0, 0
		return func() (int, bool) {
			num, next = doSomething(next), next+1
			return num, num < size
		}
	}

	for i := 0; i < b.N; i++ {
		itr := iter()
		for num, ok := itr(); ok; num, ok = itr() {
			_ = num
		}
	}
}

type iterator int

func (i *iterator) Next() int {
	out := int(*i)
	*i = iterator(doSomething(int(*i + 1)))
	return out
}

func (i iterator) Done() bool {
	return int(i) >= size
}

func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		itr := iterator(0)
		for num := itr.Next(); !itr.Done(); num = itr.Next() {
			_ = num
		}
	}
}
