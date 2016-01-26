package conmap

import "sync"

func ExampleConcurrentMap() {
	_ = ConcurrentMap{map[string]string{}, &sync.RWMutex{}}
}
